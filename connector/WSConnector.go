package connector

import (
	"bytes"
	"github.com/gorilla/websocket"
	"gitlab.com/adoontheway/goserver/codec"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: time.Duration(3) * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true //FIXME : remove after test
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

var buf []byte

type wsConnector struct {
	addr       string
	codec      codec.ICodec
	readBuf    []byte
	writeBuf   []byte
	sessionMgr *SessionManager
}

func NewWsconnector(addr string) IConnector {
	return &wsConnector{
		addr: addr,
		sessionMgr:&SessionManager{
			sessionMap:make(map[uint64]ISession),
		},
	}
}

func (wc *wsConnector) serveWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatalf("Ws serve errror:%s", err)
		return
	}

	//todo init session or not, need receive login protocol
	//todo add channel to sub&pub&broadcast
	//sess := wc.sessionMgr.NewSession()
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
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error:%v", err)
			}
			break
		}
		if err = conn.WriteMessage(websocket.BinaryMessage, buf); err != nil {

		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		log.Printf("Message Type:%d", messageType)
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

func (wc *wsConnector) SetCodec(codec codec.ICodec) {
	wc.codec = codec
}

func (wc *wsConnector) AddHandler(name string, handler HandleFunc) {
	//todo
}
