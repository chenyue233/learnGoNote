package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("监听失败，错误：", err)
		return
	}
	for {
		conn, err := listen.Accept() //建立连接
		if err != nil {
			fmt.Println("建立连接失败，错误：", err)
			continue
		}
		go process(conn) //启动一个goroutine处理连接
	}
}
