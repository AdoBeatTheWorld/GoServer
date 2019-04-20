package connector

import (
	"github.com/golang/protobuf/proto"
	"reflect"
)

type Header struct {
	PackLen   uint16
	CRC       uint16
	Version   uint16
	Sign      uint16
	MainId    uint8
	SubId     uint8
	EncryType uint8
	Back      uint8
	RequestId uint32
	DataSize  uint16
	Data      []byte
}

type MessageHandler func(msgid uint16, msg interface{})

type MessageInfo struct {
	msgType    reflect.Type
	msgHandler MessageHandler
}

var msg_map = make(map[uint16]MessageInfo)

func RegisterProtocol(msgid uint16, msg interface{}, handler MessageHandler) {
	var info MessageInfo
	info.msgType = reflect.TypeOf(msg.(proto.Message))
	info.msgHandler = handler
	msg_map[msgid] = info
}
