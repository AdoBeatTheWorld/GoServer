package server

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
	"reflect"
	"sync"
)

type IProtocolManager interface {
	RegisterProto(mainid uint8, subid uint8, msg interface{}, handler ProtoHandler) error
	GetMessageType(mainid uint8, subid uint8) (reflect.Type, error)
	Handle(mainid uint8, subid uint8, data []byte, c *Client) error
}

type protocolManager struct {
	m           sync.RWMutex
	protoMap    map[uint16]reflect.Type
	protoHander map[uint16]ProtoHandler
}

type ProtoHandler func(msg proto.Message) (result interface{}, err error)

func NewProtocolManager() IProtocolManager {
	return &protocolManager{
		protoMap:    make(map[uint16]reflect.Type),
		protoHander: make(map[uint16]ProtoHandler),
	}
}
func (pm *protocolManager) RegisterProto(mainid uint8, subid uint8, msg interface{}, handler ProtoHandler) error {
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

func (pm *protocolManager) Handle(mainid uint8, subid uint8, data []byte, c *Client) error {
	protoId := uint16(mainid << 4 & subid)
	msgType := pm.protoMap[protoId]
	handler := pm.protoHander[protoId]
	if msgType == nil {
		log.Fatalf("can not found message type for[mainid:%d,subid:%d]\n", mainid, subid)
		return errors.New(fmt.Sprintf("Message Type Not Found for [mainid:%d,subid:%d]", mainid, subid))
	}
	if handler == nil {
		log.Fatalf("can not found handler for[mainid:%d,subid:%d]\n", mainid, subid)
		return errors.New(fmt.Sprintf("Handler Not Found for [mainid:%d,subid:%d]", mainid, subid))
	}
	msg := reflect.New(msgType.Elem()).Interface()
	err := proto.Unmarshal(data, msg.(proto.Message))
	if err != nil {
		log.Fatal(fmt.Sprintf("Unmarshal Error:%s,[%d,%d]", err, mainid, subid))
	}
	_, err = handler(msg.(proto.Message))
	return nil
}
