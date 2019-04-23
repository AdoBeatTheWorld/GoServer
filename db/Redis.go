package db

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

var conn redis.Conn

func InitRedis() {
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
}

func Get(key string) {
	conn.Send("GET", key)
	conn.Flush()
	rec, err := conn.Receive()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(rec.([]byte)))
}

func Set() {

}
