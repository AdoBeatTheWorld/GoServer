package main

import (
	"flag"
	"gitlab.com/adoontheway/goserver/server"
	"os"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()

	server.StartLog()
	//server.SendRedis()

	s := server.NewServer(*addr)
	err := s.Start()

	if err != nil {
		os.Exit(1)
	}
	defer s.Stop()
}
