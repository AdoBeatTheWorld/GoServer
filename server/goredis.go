package server

import (
	"flag"
	"github.com/gomodule/redigo/redis"
	"log"
	"time"
)

var redisaddr = flag.String("rdsaddr","192.168.2.127:7001","redis address")

var RedisPool redis.Pool

func init()  {
	RedisPool = StartRedisPool(*redisaddr,"test5566","0")
}

func StartRedisPool(addr string, pass string, dbname string) redis.Pool {
	return redis.Pool{
		IdleTimeout:240*time.Second,
		MaxIdle:3,
		//Pool需要设置Dial或者DialContext，当2者都设置了的情况下，优先使用DialContext
		Dial: func() (conn redis.Conn, e error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				log.Panic("Redis dial error:", err)
				return nil, err
			}
			//密码授权
			//if _, err := c.Do("AUTH", pass); err != nil {
			//	c.Close()
			//	log.Panic("Redis auth error:", err)
			//	return nil, err
			//}
			//选择使用的数据库
			if _, err := c.Do("SELECT", dbname); err != nil {
				c.Close()
				log.Panic("Redis Select DB Error:",err)
				return nil, err
			}
			return c, nil
		},
		//在空闲连接返回到应用前检查他的健康情况
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_,err := c.Do("PING")
			return err
		},
	}
}

func SendRedis()  {
	conn := RedisPool.Get()
	err := conn.Send("SET", "A", "B")
	conn.Flush()
	if err != nil {
		log.Println("Redis Error:",err)
		return
	}
	result, err := conn.Receive()
	if err != nil {
		log.Println("Redis Send Error:",err)
		return
	}
	log.Printf("Redis Result:%+v",result)

	err = conn.Send("GET", "A")
	conn.Flush()
	if err != nil {
		log.Println("Redis Error:",err)
		return
	}
	result, err = conn.Receive()
	if err != nil {
		log.Println("Redis Send Error:",err)
		return
	}
	log.Printf("Redis Result:%+v",result)
}


