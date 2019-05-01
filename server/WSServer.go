package server

import (
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"gitlab.com/adoontheway/goserver/model"
	"log"
	"math/rand"
	"net/http"
)

type WSServer struct {
	clients  map[uint32]*Client
	sessions map[string]*model.UserInfo
	max      uint32
	current  uint32
	addr     string
}

func (s *WSServer) NewClient(conn *websocket.Conn) (*Client, error) {
	c := Client{
		conn: conn,
		s:    s,
	}
	s.clients[c.cid] = &c
	return &c, nil
}

func (s *WSServer) RemoveClient(c *Client) {
	if s.clients[c.cid] != nil {
		delete(s.clients, c.cid)
	}
}

func (s *WSServer) GenSession(account string) string {
	session, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("Gen UUID encounted:%s \n", err)
	}
	s.sessions[session.String()] = &model.UserInfo{
		Account: account,
		Name:    account,
		Score:   rand.Int63n(99999),
	}

	return session.String()
}

func (s *WSServer) Start() error {
	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/login/:account", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		Login(s, w, r, p)
	})
	router.GET("/ws", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ServeWS(s, w, r, p)
	})

	err := http.ListenAndServe(s.addr, router)
	if err != nil {
		log.Fatalln("Err occurred when start listening:", err)
		return err
	}
	log.Printf("Http is listening on:%s", s.addr)
	return nil
}

func (s *WSServer) Stop() {

}

func NewServer(addr string) *WSServer {
	return &WSServer{
		clients:  make(map[uint32]*Client),
		sessions: make(map[string]*model.UserInfo),
		//max:     maxConn,
		current: 0,
		addr:    addr,
	}
}
