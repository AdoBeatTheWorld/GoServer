package main

import (
	"gitlab.com/adoontheway/goserver/codec"
	"gitlab.com/adoontheway/goserver/connector"
	"os"
)

func init() {
	codec.InitCodec()
}

func main() {
	codec.RegisterCodec(codec.CodecJson, codec.NewJsonCode())
	conn := connector.NewWsconnector(":8080")
	err := conn.Start()
	if err != nil {
		conn.Stop()
		os.Exit(1)
	}
}
