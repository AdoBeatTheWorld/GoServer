package test

import (
	"bytes"
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"gitlab.com/adoontheway/goserver/protocol"
	"log"
	"net/http"
	"testing"
)

type CHeader struct {
	Len       uint16
	MainId    uint8
	SubId     uint8
	EncryType uint8
	data      []byte
}

func TestConnection(t *testing.T) {
	dialer := websocket.Dialer{}
	header := http.Header{}
	header.Add("host", "127.0.0.1")
	header.Add("name", "test")
	conn, _, err := dialer.Dial("ws://127.0.0.1:8080/ws", header)
	if err != nil {
		log.Fatalln(err)
	}
	login := &protocol.Login{
		Session: "Incrediable Hulk.",
	}

	data, err := proto.Marshal(login)

	cheader := &CHeader{
		MainId:    1,
		SubId:     1,
		EncryType: 0,
		Len:       uint16(5 + len(data)),
		data:      data,
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, cheader.Len)       //len
	binary.Write(buf, binary.LittleEndian, cheader.MainId)    //mainid
	binary.Write(buf, binary.LittleEndian, cheader.SubId)     //subid
	binary.Write(buf, binary.LittleEndian, cheader.EncryType) //encry type
	binary.Write(buf, binary.LittleEndian, cheader.data)      //data

	conn.WriteMessage(websocket.BinaryMessage, buf.Bytes())
}
