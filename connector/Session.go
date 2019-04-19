package connector

import (
	"fmt"
	"log"
	"net"
)

type ISession interface {
	Start()
	Stop()
	Dispose()
	GetId() uint64
	String() string
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

func (s *session) Start() {
	//todo
	log.Printf("Session Start:%d", s.ID)
}

func (s *session) Stop() {
	//todo
	log.Printf("Session Stop:%d", s.ID)
}

func (s *session) Dispose() {
	//todo
	log.Printf("Session Displose:%d", s.ID)
}

func (s *session) GetId() uint64 {
	return s.ID
}

func (s *session) String() string {
	return fmt.Sprintf("Session(id:%d,userid:%d,addr:%s)", s.ID, s.UserId, s.conn.RemoteAddr().String())
}
