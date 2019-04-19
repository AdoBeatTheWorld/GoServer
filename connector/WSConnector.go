package connector

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: time.Duration(3) * time.Second,
}

const MaxMessageSize int64 = 1024

type wsConnector struct {
	addr string
}

func NewWsconnector(addr string) IConnector {
	return &wsConnector{
		addr: addr,
	}
}

func serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	defer func() {
		conn.Close()
	}()

	conn.SetReadLimit(MaxMessageSize)

	if err != nil {
		log.Fatalf("Ws serve errror:%s", err)
		return
	}
	//todo init session or not, need receive login protocol
	//todo add channel to sub&pub&broadcast
}

func (wc *wsConnector) Start() error {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(w, r)
	})
	err := http.ListenAndServe(wc.addr, nil)
	return err
}

func (wc *wsConnector) Stop() {
	//todo
}
