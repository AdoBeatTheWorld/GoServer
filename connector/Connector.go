package connector

import (
	"bufio"
	"gitlab.com/adoontheway/goserver/codec"
)

type HandleFunc func(writer bufio.ReadWriter)

type IConnector interface {
	Start() error
	Stop()
	SetCodec(codec codec.ICodec)
	AddHandler(name string, handler HandleFunc)
}
