package client

import (
	"github.com/gorilla/websocket"
)

type ClientManager struct {
	clientMap map[uint32]*Client
	max uint32
	current uint32
}

type Client struct {
	conn *websocket.Conn
	cid uint32
	uid uint32
	session string
}

func (c *Client) Start()  {
	go c.startWriteLoop()
	go c.startReadLoop()
}

func (c *Client) startReadLoop()  {
	
}

func (c *Client) startWriteLoop()  {
	
}

func (cm *ClientManager) NewClient(conn *websocket.Conn) (*Client,error) {
	//num := cm.GetNumClients()
	//if num >= cm.max {
	//	return nil, errors.New("Connection Over Limit.")
	//}
	//fmt.Println(num)
	cm.current++;
	c := Client{
		conn:conn,
		//cid: num + 1,
	}
	cm.clientMap[c.cid] = &c
	return &c,nil
}

func (cm *ClientManager) GetMax() uint32 {
	return cm.max
}

//func (cm *ClientManager) GetNumClients() uint32 {
//	return uint32(len(cm.clientMap))
//}

func NewClientManager(maxConn uint32) *ClientManager {
	return &ClientManager{
		clientMap:make(map[uint32]*Client),
		max:maxConn,
		current : 0,
	}
}