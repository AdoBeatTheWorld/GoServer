package connector

import (
	"bytes"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: time.Duration(3) * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	maxMessageSize int64 = 1024
	pongWait             = 60 * time.Second
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

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

	if err != nil {
		log.Fatalf("Ws serve errror:%s", err)
		return
	}
	//todo init session or not, need receive login protocol
	//todo add channel to sub&pub&broadcast
	go startReadLoop(conn)
	go startWriteLoop(conn)
}

func startReadLoop(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	conn.SetReadLimit(maxMessageSize)
	conn.SetReadDeadline(time.Now().Add(pongWait))
	//conn.PongHandler(func(string) error {conn.SetReadDeadline(time.Now().Add(pongWait));	return nil})
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error:%v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		log.Printf("Recieved message:%s", message)
	}
}

func startWriteLoop(conn *websocket.Conn) {

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
