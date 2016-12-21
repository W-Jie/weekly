package main

import (
	"github.com/garyburd/redigo/redis"
)

// 转存到redis
func toredis(server string, r string) {
	conn, err := redis.Dial("tcp", server)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Do("lpush", "weeklyreport", r)
	if err != nil {
		panic(err)
	}
	//fmt.Println(v)
}
