package rpc

import (
	"fmt"
	"net"
	"time"

	"black-hole.com/modules/pool"
)

var Pool pool.Pool

func init() {
	//factory 创建连接的方法
	factory := func() (interface{}, error) { return net.Dial("tcp", addr) }

	//close 关闭连接的方法
	close := func(v interface{}) error { return v.(net.Conn).Close() }

	//创建一个连接池： 初始化2，最大连接5，空闲连接数是4
	poolConfig := &pool.Config{
		InitialCap: 2,
		MaxIdle:    4,
		MaxCap:     5,
		Factory:    factory,
		Close:      close,
		//连接最大空闲时间，超过该时间的连接 将会关闭，可避免空闲时连接EOF，自动失效的问题
		IdleTimeout: 15 * time.Second,
	}
	p, err := pool.NewChannelPool(poolConfig)
	if err != nil {
		fmt.Println("err=", err)
	}
}
