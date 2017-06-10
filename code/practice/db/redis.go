package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

// 常用的redis驱动
// https://github.com/garyburd/redigo
// https://github.com/hoisie/redis

func main() {

	// connect redis database
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// set a value
	isSuccess, err := redis.String(conn.Do("SET", "name", "allen"))
	if err != nil {
		panic(err)
	}

	if isSuccess == "OK" {
		fmt.Println("success")
	} else {
		fmt.Println("set name failed!")
	}

}
