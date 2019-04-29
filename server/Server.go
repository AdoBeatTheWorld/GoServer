package server

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
)

type Server struct {
	clients map[uint32]*Client
	max     uint32
	current uint32
}

func (s *Server) NewClient(conn *websocket.Conn) (*Client, error) {
	c := Client{
		conn: conn,
		s:    s,
	}
	s.clients[c.cid] = &c
	return &c, nil
}

func (s *Server) GenSession(account string) string {
	session, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Gen UUID encounted:%s \n", err)
	}
	//s.sessions[session.String()] = &model.UserInfo{
	//	Account:account,
	//	Name:account,
	//	Score:rand.Int63n(99999),
	//}

	return session.String()
}

func NewServer(maxConn uint32) *Server {
	return &Server{
		clients: make(map[uint32]*Client),
		//sessions: make(map[string]*model.UserInfo),
		max:     maxConn,
		current: 0,
	}
}
