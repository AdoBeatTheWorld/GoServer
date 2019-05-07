package server

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

var redisPool redis.Pool

func StartRedisPool(addr string, pass string, dbname string) {
	redisPool = redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Println("Redis dial error:", err)
				return nil, err
			}
			if _, err := c.Do("AUTH", pass); err != nil {
				c.Close()
				log.Println("Redis auth error:", err)
				return nil, err
			}
			if _, err := c.Do("SELECT", dbname); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}
}
