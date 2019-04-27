package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
	"gitlab.com/adoontheway/goserver/connector"
	"gitlab.com/adoontheway/goserver/proto"
	"log"
	"math/rand"
	"net/http"
	"runtime"
	"strings"
)

var localIp string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	httpserver := connector.NewHttpServer(":8081")
	httpserver.AddHandler("/", index, http.MethodGet)
	httpserver.AddHandler("/login", login, http.MethodGet)
	httpserver.Serve()
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
