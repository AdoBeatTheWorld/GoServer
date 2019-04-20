package connector

import "gitlab.com/adoontheway/goserver/codec"

type IConnector interface {
	Start() error
	Stop()
	SetCodec(codec codec.ICodec)
}
