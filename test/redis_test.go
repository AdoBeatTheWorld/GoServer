package test

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"testing"
	"time"
)

func TestRedisPool(t *testing.T) {

}

func TestRedis(t *testing.T) {
	readTimeOut := redis.DialReadTimeout(time.Second * 3)
	connectTimeOut := redis.DialConnectTimeout(time.Second * 10)
	writeTimeOut := redis.DialWriteTimeout(time.Second * 3)
	conn, err := redis.Dial("tcp", "192.168.2.127:7001", readTimeOut, connectTimeOut, writeTimeOut)
	if err != nil {
		fmt.Println("Dial Failed:", err)
		t.FailNow()
	}
	err = conn.Send("SET", "foo", "bar")
	if err != nil {
		fmt.Println("Set Failed:", err)
		t.FailNow()
	}
}
