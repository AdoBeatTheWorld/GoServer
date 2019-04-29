package server

import (
	"gitlab.com/adoontheway/goserver/protocol"
	"gitlab.com/adoontheway/goserver/server/handler"
	"log"
)

var ProtoMgr IProtocolManager

func init() {
	ProtoMgr = NewProtocolManager()
	err := ProtoMgr.RegisterProto(HALL, LOGIN, &protocol.Login{}, handler.LoginHandler)
	if err != nil {
		log.Fatalln(err)
	}
}
