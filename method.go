package qn

import "strings"

type ProtocolType byte

const (
	Http ProtocolType = 1 << iota
	EvenBus
	Websocket
)

var pts = map[string]ProtocolType{
	"http": Http,
	"eb":   EvenBus,
	"ws":   Websocket,
}

type Method string

func (m Method) GetProtocolType() ProtocolType {
	s := strings.Split(string(m), "_")
	return pts[s[0]]
}

func (m Method) GetAction() string {
	s := strings.Split(string(m), "_")
	if len(s) > 1 {
		return strings.ToUpper(s[1])
	}
	return ""
}

const (
	HttpPost   Method = "http_post"
	HttpGet    Method = "http_get"
	HttpPut    Method = "http_put"
	HttpDelete Method = "http_delete"
	EB         Method = "eb"
	WS         Method = "ws"
)
