package connector

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

type MyHttpServer interface {
	Serve() error
	AddHandler(path string, handler httprouter.Handle, method string)
}

type httpServer struct {
	addr   string
	router *httprouter.Router
}

func (hs *httpServer) Serve() error {
	err := http.ListenAndServe(hs.addr, hs.router)
	if err != nil {
		log.Fatalf("Http Server start encountered error : %s", err)
		return err
	}
	log.Printf("Http Server is listening on:%s", hs.addr)
	return nil
}

func (hs *httpServer) AddHandler(path string, handler httprouter.Handle, method string) {
	if method == "GET" {
		hs.router.GET(path, handler)
	} else if method == "POST" {
		hs.router.POST(path, handler)
	} else {
		log.Fatalf("Method:%s not support for now.", method)
	}
}

func NewHttpServer(addr string) MyHttpServer {
	return &httpServer{
		addr:   addr,
		router: httprouter.New(),
	}
}
