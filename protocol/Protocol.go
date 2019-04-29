package protocol

import (
	"github.com/golang/protobuf/proto"
	"reflect"
	"sync"
)

type IProtocolManager interface {
	RegisterProto(mainid uint8, subid uint8, msg *proto.Message, handler ProtoHandler) error
	GetMessageType(mainid uint8, subid uint8) (reflect.Type, error)
}

type protocolManager struct {
	m           sync.RWMutex
	protoMap    map[uint16]reflect.Type
	protoHander map[uint16]ProtoHandler
}

type ProtoHandler func(msgid uint16, msg proto.Message) (result interface{}, err error)

func NewProtocolManager() IProtocolManager {
	return &protocolManager{
		protoMap:    make(map[uint16]reflect.Type),
		protoHander: make(map[uint16]ProtoHandler),
	}
}
func (pm *protocolManager) RegisterProto(mainid uint8, subid uint8, msg *proto.Message, handler ProtoHandler) error {
	protoId := uint16(mainid << 4 & subid)
	pm.m.Lock()
	defer pm.m.Unlock()
	pm.protoMap[protoId] = reflect.TypeOf(msg)
	pm.protoHander[protoId] = handler
	return nil
}

func (pm *protocolManager) GetMessageType(mainid uint8, subid uint8) (reflect.Type, error) {
	protoId := uint16(mainid << 4 & subid)
	return pm.protoMap[protoId], nil
}
