package client

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
	"gitlab.com/adoontheway/goserver/model"
	"log"
	"math/rand"
	"time"
)

const (
	maxMessageSize = 1024
	pongWait = time.Second * 3
)

type Header struct {
	Len uint16//package length
	MainId uint8//main protocol id
	SubId uint8//sub protocol id
	EncryType uint8//encrypto type
	data []byte
}

type ClientManager struct {
	clientMap map[uint32]*Client
	sessions map[string]*model.UserInfo
	max uint32
	current uint32
}

type Client struct {
	conn *websocket.Conn
	cid uint32
	uid uint32
	session string
	manager *ClientManager
}

func (c *Client) Start()  {
	go c.startWriteLoop()
	go c.startReadLoop()
}

func (c *Client) startReadLoop()  {
	c.conn.SetReadLimit(maxMessageSize)
	err := c.conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		log.Printf("SetReadDeadLine Error:%s",err)
	}
	c.conn.SetPongHandler(func(appData string) error {
		err = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			log.Printf("SetReadDeadLine Error:%s",err)
		}
		return nil
	})
	//for {
	//	mType,data,err := c.conn.ReadMessage()
	//	if err != nil {
	//		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
	//			log.Print("error:%v",err)
	//		}
	//		break
	//	}
	//	msg := message.Message{}
	//	err := proto.Unmarshal(data, &msg)
	//}
}

func (c *Client) startWriteLoop()  {
	
}

func (cm *ClientManager) NewClient(conn *websocket.Conn) (*Client,error) {
	num := cm.GetNumClients()
	if num >= cm.max {
		return nil, errors.New("Connection Over Limit.")
	}
	cm.current++;
	c := Client{
		conn:conn,
		cid: cm.current,
		manager:cm,
	}
	cm.clientMap[c.cid] = &c
	return &c,nil
}

func (cm *ClientManager) GetMax() uint32 {
	return cm.max
}

func (cm *ClientManager) GetNumClients() uint32 {
	return uint32(len(cm.clientMap))
}

func (cm *ClientManager) GenSession(account string) string {
	session,err := uuid.NewV4()
	if err!=nil {
		log.Fatalf("Gen UUID encounted:%s \n",err)
	}
	cm.sessions[session.String()] = &model.UserInfo{
		Account:account,
		Name:account,
		Score:rand.Int63n(99999),
	}

	return session.String()
}

func NewClientManager(maxConn uint32) *ClientManager {
	return &ClientManager{
		clientMap:make(map[uint32]*Client),
		sessions: make(map[string]*model.UserInfo),
		max:maxConn,
		current : 0,
	}
}