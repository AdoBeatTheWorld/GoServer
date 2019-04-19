package main

import (
	"gitlab.com/adoontheway/goserver/codec"
	"gitlab.com/adoontheway/goserver/connector"
)

func init() {
	codec.InitCodec()

}

func main() {
	codec.RegisterCodec(codec.NewJsonCode())
	connector.NewTcpConnector(":8080").Start()
}
