package main

import (
	"flag"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/adoontheway/goserver/server"
	"log"
	"net/http"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var addr = flag.String("add", ":8080", "http service address")

func main() {
	flag.Parse()

	server.StartLog()

	s := server.NewServer(100)

	router := httprouter.New()
	router.GET("/", server.Index)
	router.POST("/login/:account", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		server.Login(s, w, r, p)
	})
	router.GET("/ws", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		server.ServeWS(s, w, r, p)
	})

	err := http.ListenAndServe(*addr, router)
	checkErr(err)
	log.Printf("Http is listening on:%s", *addr)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
