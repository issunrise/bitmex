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
	ws          *websocket.Conn
	readSignal  chan *Response
	writeSignal chan interface{}
	doneSignal  chan bool
	err         error
	heartCnt    int
}

func NewDataSignal(size ...uint64) chan *Response {
	if len(size) > 0 {
		return make(chan *Response, size[0])
	}
	return make(chan *Response, RBUFFER)
}

func NewWsClient(readSignal chan *Response, dialer proxy.Dialer) (*WebSocketClient, error) {
	client := websocket.DefaultDialer
	client.ReadBufferSize = 10240
	client.ReadBufferSize = 10240
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
		ws:          ws,
		readSignal:  readSignal,
		writeSignal: make(chan interface{}, WBUFFER),
		doneSignal:  make(chan bool, 10),
	}, nil
}

func (c *WebSocketClient) ListenWrite() {
	heart := time.NewTicker(time.Second * 6)
	defer func() {
		recover()
		heart.Stop()
		close(c.writeSignal)
		log.Println("ListenWrite closed")
	}()
	for {
		select {
		case <-c.doneSignal:
			return
		case <-heart.C:
			c.Write("ping")
			if c.heartCnt++; c.heartCnt > 3 {
				c.err = errors.New(fmt.Sprintf("Webscoket connection timeout. heart cnt is %d", c.heartCnt))
				close(c.doneSignal)
				return
			}
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
		recover()
		log.Println("ListenRead closed")
	}()
	for {
		select {
		case <-c.doneSignal:
			return
		default:
			var rep Response
			_, data, err := c.ws.ReadMessage()
			if err != nil {
				c.err = errors.New(fmt.Sprintf("Webscoket read error: %v", err))
				close(c.doneSignal)
				return
			}
			c.heartCnt = 0
			if string(data) == "pong" {
				continue
			}
			// log.Println("bitmex websocket data:", string(data))
			if err := json.Unmarshal(data, &rep); err != nil {
				log.Printf("Webscoket unmarshal error:%v,data:%v", err, string(data))
			} else {
				c.readSignal <- &rep
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

// func (c *WebSocketClient) ListenHeart() {
// 	defer func() {
// 		c.exitCnt++
// 		c.heart.timer.Stop()
// 		log.Println("ListenHeart closed")
// 	}()
// 	for {
// 		select {
// 		case <-c.doneSignal:
// 			return
// 		case <-c.heart.timer.C:
// 			c.Write("ping")
// 			c.heart.cnt++
// 			// log.Println("Send ping times:", c.heart.cnt)
// 			if c.heart.cnt > 3 {
// 				c.err = errors.New(fmt.Sprintf("Webscoket connection timeout. heart cnt is %d", c.heart.cnt))
// 				close(c.doneSignal)
// 				return
// 			}
// 		}
// 	}
// }

func (c *WebSocketClient) Run() error {
	go c.ListenWrite()
	c.ListenRead()
	return c.err
}

func (c *WebSocketClient) Close() error {
	defer func() {
		recover()
	}()
	close(c.doneSignal)
	c.err = errors.New("close conn.")
	return c.ws.Close()
}

func (c *WebSocketClient) Write(in interface{}) {
	defer func() {
		recover()
	}()
	select {
	case c.writeSignal <- in:
	}

}

func (c *WebSocketClient) Subscribe(args []string) {
	c.Write(Msg{"subscribe", args})
}
