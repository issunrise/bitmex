package ws

type Msg struct {
	Op   string   `json:"op,omitempty"`
	Args []string `json:"args,omitempty"`
}

type Response struct {
	Success   bool        `json:"success,omitempty"`
	Subscribe string      `json:"subscribe,omitempty"`
	Request   interface{} `json:"request,omitempty"`
	Table     string      `json:"table,omitempty"`
	Action    string      `json:"action,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

type BitmexDepth struct {
	Id     int64   `json:"id,omitempty"`
	Symbol string  `json:"symbol,omitempty"`
	Side   string  `json:"side,omitempty"`
	Price  float64 `json:"price,omitempty"`
	Size   float64 `json:"size,omitempty"`
}
