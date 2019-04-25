package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"gitlab.com/adoontheway/goserver/codec"
	"gitlab.com/adoontheway/goserver/connector"
	"gitlab.com/adoontheway/goserver/proto"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"
)

var localIp string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	codec.InitCodec()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops:" + err.Error())
		os.Exit(1)
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			localIp = ipnet.IP.String()
			log.Printf(" Local Ip : %s", localIp)
		}
	}
}

func main() {
	httpserver := connector.NewHttpServer(":8081")
	httpserver.AddHandler("/", index, http.MethodGet)
	httpserver.AddHandler("/login", login, http.MethodGet)
	httpserver.Serve()

	codec.RegisterCodec(codec.CodecJson, codec.NewJsonCode())
	wsserver := connector.NewWsconnector(":8080")
	err := wsserver.Start()
	if err != nil {
		wsserver.Stop()
		os.Exit(1)
	}
}

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "Welcome....")
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var account string //= ps.ByName("account")
	account = r.URL.RawQuery
	if len(account) != 0 {
		strs := strings.Split(account, "=")
		if strs[0] == "account" {
			account = strs[1]
		}
	}
	log.Printf("Account:%s is logging in", account)
	uid, err := uuid.FromString(account)
	if err != nil {
		fmt.Fprintf(w, "login error:%s", err)
		return
	}
	loginResult := &proto.LoginResult{
		Uid:        uid.String(),
		Account:    account,
		Score:      rand.Int63n(100000),
		GameServer: fmt.Sprintf("%s:%s", localIp, 8080),
	}
	data, err := json.Marshal(loginResult)
	if err != nil {
		fmt.Fprintf(w, "Error encounted:%s", err)
		return
	}
	fmt.Sprint(w, string(data))
}
