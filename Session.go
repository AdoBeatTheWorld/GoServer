package main

import (
	"log"
	"net"
)

type ISession interface {
	Start()
	Stop()
	Dispose()
}

type session struct {
	ID     uint64
	UserId uint64
	conn   net.Conn
}

type sessionManager struct {
	sessionMap map[uint64]ISession
}

var manager sessionManager

var flag int32 = 0

func Start() error {
	manager = sessionManager{
		sessionMap: make(map[uint64]ISession),
	}
	return nil
}

func NewSession(id uint64, userid uint64, conn net.Conn) (ISession, error) {
	s := &session{
		ID:     id,
		UserId: userid,
		conn:   conn,
	}
	manager.sessionMap[id] = s
	return s, nil
}

func (s session) Start() {
	//todo
	log.Printf("Session Start:%d", s.ID)
}

func (s session) Stop() {
	//todo
	log.Printf("Session Stop:%d", s.ID)
}

func (s session) Dispose() {
	//todo
	log.Printf("Session Displose:%d", s.ID)
}
