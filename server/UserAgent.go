package server

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
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
	s       *WSServer
}

func (c *Client) Start() {
	//go c.startWriteLoop()
	go c.startReadLoop()
}

func (c *Client) Send(mainId uint8, subId uint8, message proto.Message) {
	buf, err := proto.Marshal(message)
	if err != nil {
		log.Fatalln("proto marshall error:", err)
	}
	len := len(buf) + 5
	dataBuf := bytes.NewBuffer([]byte{})
	err = binary.Write(dataBuf, binary.LittleEndian, len)
	if err != nil {
		log.Println("Error occurred when write stream:", err)
	}
	err = binary.Write(dataBuf, binary.LittleEndian, mainId)
	if err != nil {
		log.Println("Error occurred when write stream:", err)
	}
	err = binary.Write(dataBuf, binary.LittleEndian, subId)
	if err != nil {
		log.Println("Error occurred when write stream:", err)
	}
	err = binary.Write(dataBuf, binary.LittleEndian, 0)
	if err != nil {
		log.Println("Error occurred when write stream:", err)
	}
	err = binary.Write(dataBuf, binary.LittleEndian, buf)
	if err != nil {
		log.Println("Error occurred when write stream:", err)
	}
	err = c.conn.WriteMessage(websocket.BinaryMessage, dataBuf.Bytes())
	if err != nil {
		log.Println("Error occurred when write stream:", err)
	}
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
		reader := bytes.NewReader(data)
		//check legal

		//parse
		header := &Header{}
		err = binary.Read(reader, binary.LittleEndian, &header.Len)
		if int(header.Len) > len(data) {
			log.Println("Illegal packet.")
			c.Stop()
			return
		}
		err = binary.Read(reader, binary.LittleEndian, &header.MainId)
		err = binary.Read(reader, binary.LittleEndian, &header.SubId)
		err = binary.Read(reader, binary.LittleEndian, &header.EncryType)
		err = binary.Read(reader, binary.LittleEndian, &header.data)
		ProtoMgr.Handle(header.MainId, header.SubId, header.data, c)
	}
}

func (c *Client) startWriteLoop() {

}

func (c *Client) Stop() {
	c.s.RemoveClient(c)
	err := c.conn.Close()
	if err != nil {
		log.Printf("Close client error:%s", err)
	}
}
