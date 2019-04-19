package connector

import (
	"github.com/astaxie/beego/logs"
	"log"
	"net"
)

type tcpConnector struct {
	Addr     string
	readbuf  [1024]byte
	writebuf [1024]byte
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

func (tc *tcpConnector) Stop() {
	//todo
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

func handleConn(conn net.Conn) error {
	defer conn.Close()
	for {
		var buf = make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		log.Printf("Read %d bytes, content is %s\n", n, string(buf[:n]))
	}
}
