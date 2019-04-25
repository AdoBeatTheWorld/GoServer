package connector

import (
	"bufio"
	"github.com/astaxie/beego/logs"
	"gitlab.com/adoontheway/goserver/codec"
	"log"
	"net"
	"sync"
)

type tcpConnector struct {
	m       sync.RWMutex
	Addr    string
	handler map[string]HandleFunc
	codec   codec.ICodec
}

func NewTcpConnector(addr string) IConnector {
	return &tcpConnector{
		Addr: addr,
	}
}

func (tc *tcpConnector) Start() error {
	logs.Info("Tcp Connector is Listening...")
	listener, err := net.Listen("tcp", tc.Addr)
	if err != nil {
		return err
	}
	go run(listener)

	return nil
}

func (tc *tcpConnector) SetCodec(codec codec.ICodec) {
	//todo
	tc.codec = codec
}

func (tc *tcpConnector) Stop() {
	//todo
}

func (tc *tcpConnector) AddHandler(name string, handler HandleFunc) {
	//todo
	tc.m.Lock()
	tc.handler[name] = handler
	tc.m.Unlock()
}

func run(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			logs.Info(err)
			break
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	rw := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	defer conn.Close()

	for {
		var newBytes = make([]byte, 1024)
		n, err := rw.Read(newBytes) //conn.Read(buf)
		if err != nil {
			log.Fatalf("handleConn:%s", err)
			break
		}

		if n < 1024 {

		}

		log.Printf("Read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
