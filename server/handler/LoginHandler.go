package handler

import (
	"github.com/golang/protobuf/proto"
	"log"
)

func LoginHandler(msg proto.Message) (result interface{}, err error) {
	log.Printf("LoginHandler:%x", msg.String())
	return nil, nil
}
