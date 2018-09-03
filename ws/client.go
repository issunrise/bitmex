package ws

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	// "golang.org/x/net/websocket"
	"github.com/gorilla/websocket"
	"golang.org/x/net/proxy"
)

const (
	ORIGIN  = "http://localhost"
	BITWS   = "wss://www.bitmex.com/realtime"
	WBUFFER = 1024
	RBUFFER = 1024
)

type WebSocketClient struct {
	heart       *Heart
	ws          *websocket.Conn
	readSignal  chan *Response
	writeSignal chan interface{}
	doneSignal  chan bool
	err         error
	exitCnt     int
}

func NewDataSignal() chan *Response {
	return make(chan *Response, RBUFFER)
}

func NewWsClient(data chan *Response, dialer proxy.Dialer) (*WebSocketClient, error) {

	// var ws *websocket.Conn
	// var err error
	// if dialer != nil {
	// 	link, _ := url.Parse(BITWS)
	// 	ws, _, err = websocket.NewClient(agent, link, nil, 1024, 1024)
	// 	if err != nil {
	// 		log.Println("NewClient", err)
	// 		return nil, err
	// 	}
	// } else {
	// 	var d = websocket.Dialer{
	// 		Subprotocols:    []string{"p1", "p2"},
	// 		ReadBufferSize:  1024,
	// 		WriteBufferSize: 1024,
	// 		Proxy:           http.ProxyFromEnvironment,
	// 	}
	// 	d.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}

	// 	ws, _, err = d.Dial(BITWS, nil)
	// 	if err != nil {
	// 		log.Println("Dial", err)
	// 		return nil, err
	// 	}
	// }
	client := websocket.DefaultDialer
	client.ReadBufferSize = 10240
	client.ReadBufferSize = 10240
	log.Println("new-----------------")
	// client := &websocket.Dialer{
	// 	Subprotocols:    []string{"p1", "p2"},
	// 	ReadBufferSize:  1024,
	// 	WriteBufferSize: 1024,
	// 	// Proxy:           http.ProxyFromEnvironment,
	// 	TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	// }

	if dialer != nil {
		client.NetDial = dialer.Dial
	}

	ws, _, err := client.Dial(BITWS, nil)
	if err != nil {
		return nil, err
	}
	return &WebSocketClient{
		NewHeart(),
		ws,
		data,
		make(chan interface{}, WBUFFER),
		make(chan bool, 20),
		nil,
		0,
	}, err
}

func (c *WebSocketClient) ListenWrite() {
	defer func() {
		c.exitCnt++
		log.Println("ListenWrite closed")
	}()
	for {
		select {
		case <-c.doneSignal:
			c.doneSignal <- true
			return
		case data := <-c.writeSignal:
			// log.Println("writeSignal:", data)
			msg, _ := json.Marshal(data)
			if err := c.ws.WriteMessage(websocket.TextMessage, msg); err != nil {
				log.Printf("Websocket send data error:%v,data:%v", err, data)
			}
			// if err := websocket.JSON.Send(c.ws, data); err != nil {
			// 	log.Printf("Websocket send data error:%v,data:%v", err, data)
			// }
		}
	}

}

func (c *WebSocketClient) ListenRead() {
	defer func() {
		c.exitCnt++
		log.Println("ListenRead closed")
	}()
	for {
		select {
		case <-c.doneSignal:
			log.Println("old data close")
			c.doneSignal <- true
			return
		default:
			var rep Response
			_, msg, err := c.ws.ReadMessage()
			if err != nil {
				log.Printf("Webscoket read error:%v", err)
				c.err = err
				c.doneSignal <- true
				return
			}
			if err := json.Unmarshal(msg, &rep); err != nil {
				if string(msg) == "pong" {
					c.heart.Reset()
				} else {
					log.Printf("Webscoket unmarshal error:%v,data:%v", err, string(msg))
				}
			} else {
				c.readSignal <- &rep
				c.heart.Reset()
			}
		}
	}

	// for {
	// 	select {
	// 	case <-c.doneSignal:
	// 		c.Done("read.")
	// 		log.Println("ListenRead closed.")
	// 		return
	// 	default:
	// 		var rep Response
	// 		err := websocket.JSON.Receive(c.ws, &rep)
	// 		if err == io.EOF {
	// 			c.err = errors.New("Webscoket connection break, error:" + err.Error())
	// 			c.Done("read.")
	// 		} else if err != nil {
	// 			var msg string
	// 			if err = websocket.Message.Receive(c.ws, &msg); err != nil {
	// 				log.Printf("Webscoket read error:%v", err)
	// 			} else {
	// 				c.heart.Reset()
	// 				// log.Printf("Webscoket read heart msg:%v", msg)
	// 			}
	// 		} else {
	// 			// log.Println("data----:", rep)
	// 			c.readSignal <- &rep
	// 			c.heart.Reset()
	// 		}
	// 	}
	// }
}

func (c *WebSocketClient) ListenHeart() {
	defer func() {
		c.exitCnt++
		c.heart.timer.Stop()
		log.Println("ListenHeart closed")
	}()
	for {
		select {
		case <-c.doneSignal:
			c.doneSignal <- true
			return
		case <-c.heart.timer.C:
			c.Write("ping")
			c.heart.cnt++
			// log.Println("Send ping times:", c.heart.cnt)
			if c.heart.cnt > 3 {
				c.err = errors.New(fmt.Sprintf("Webscoket connection timeout. heart cnt is %d", c.heart.cnt))
				c.doneSignal <- true
				return
			}
		}
	}
}

func (c *WebSocketClient) Run() error {

	go c.ListenWrite()
	go c.ListenRead()
	c.ListenHeart()

	for c.exitCnt < 3 {
		time.Sleep(time.Millisecond * 200)
	}
	close(c.doneSignal)
	close(c.writeSignal)
	log.Println("bitmex websocket closed. close signal:", c.exitCnt)
	return c.err
}

func (c *WebSocketClient) Close() error {
	c.doneSignal <- true
	c.err = errors.New("close conn.")
	return c.ws.Close()
}

func (c *WebSocketClient) Write(in interface{}) {
	c.writeSignal <- in
}

func (c *WebSocketClient) Subscribe(args []string) {
	c.Write(Msg{"subscribe", args})
}
