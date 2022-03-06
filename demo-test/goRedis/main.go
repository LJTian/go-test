package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {
	Ip := "127.0.0.1"
	Port := "49153"

	// 建立链接
	c, err := redis.Dial("tcp", Ip+":"+Port)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()

	// 获取单个Key值
	data, err := c.Do("get", "name_1")
	if err != nil {
		return
	}
	Buff, err := redis.String(data, err)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Buff)

	// 获取一群Key值
	var Keys []string

	data1, err := redis.Values(c.Do("KEYS", "name*"))
	if err != nil {
		return
	}

	err = redis.ScanSlice(data1, &Keys)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(Keys)

}
