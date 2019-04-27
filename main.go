package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/adoontheway/goserver/client"
	"log"
	"net/http"
	"runtime"
)

func init()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

var addr = flag.String("add",":8080","http service address")

var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
}

func main()  {
	flag.Parse()

	clientManager := client.NewClientManager(100)

	router := httprouter.New()
	router.GET("/", Index)
	router.POST("/login",Login)
	router.GET("/ws", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		ServeWS(clientManager,w,r,p)
	})

	err := http.ListenAndServe(*addr, router)
	checkErr(err)
	log.Printf("Http is listening on:%s",*addr)
}

func ServeWS(clientManager *client.ClientManager,w http.ResponseWriter,r *http.Request, _ httprouter.Params)  {
	conn, err := upgrader.Upgrade(w,r,nil)
	checkErr(err)
	c,err := clientManager.NewClient(conn)
	checkErr(err)
	c.Start()
}

func Index(w http.ResponseWriter,r *http.Request, _ httprouter.Params)  {
	log.Print(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed.", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w,r,"./resource/index.html")
}
//return token then login through ws with token
//in production env, the token is generated by page
// this function must be depricated in production
func Login(w http.ResponseWriter,r *http.Request, _ httprouter.Params)  {
	n,err := fmt.Fprint(w,"Login...\n")
	log.Printf("Writed:%d bytes",n)
	checkErr(err)
}

func checkErr(err error)  {
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}
}

