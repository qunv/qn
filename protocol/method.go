package protocol

type ProtocolType byte

const (
	Http ProtocolType = 1 << iota
	EvenBus
	Websocket
)

type Method interface {
	GetProtocolType() ProtocolType
	GetAction() string
}
