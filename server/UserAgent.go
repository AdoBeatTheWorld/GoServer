package server

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

const (
	maxMessageSize = 1024
	pongWait       = time.Second * 3
)

type Header struct {
	Len       uint16
	MainId    uint8
	SubId     uint8
	EncryType uint8
	data      []byte
}

type Client struct {
	conn    *websocket.Conn
	send    chan []byte
	cid     uint32
	uid     uuid.UUID
	session string
	rw      bufio.ReadWriter
	s       *Server
}

func (c *Client) Start() {
	go c.startWriteLoop()
	go c.startReadLoop()
}

func (c *Client) startReadLoop() {
	c.conn.SetReadLimit(maxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		log.Printf("SetReadDeadLine Error:%s", err)
	}
	c.conn.SetPongHandler(func(appData string) error {
		err = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			log.Printf("SetReadDeadLine Error:%s", err)
		}
		return nil
	})
	for {
		_, data, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error:%v", err)
			}
			break
		}
		header := &Header{}
		reader := bytes.NewReader(data)
		//binary.Read(reader, binary.LittleEndian, &header)
		binary.Read(reader, binary.LittleEndian, &header.Len)
		header.data = make([]byte, header.Len-5)
		binary.Read(reader, binary.LittleEndian, &header.MainId)
		binary.Read(reader, binary.LittleEndian, &header.SubId)
		binary.Read(reader, binary.LittleEndian, &header.EncryType)
		binary.Read(reader, binary.LittleEndian, &header.data)

		ProtoMgr.Handle(header.MainId, header.SubId, header.data, c)
	}
}

func (c *Client) startWriteLoop() {

}
