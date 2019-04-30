package handler

import (
	"github.com/golang/protobuf/proto"
	"gitlab.com/adoontheway/goserver/protocol"
	"log"
)

func LoginHandler(msg proto.Message) (result interface{}, err error) {
	log.Printf("LoginHandler:%s", msg.(*protocol.Login).Session)
	return nil, nil
}
