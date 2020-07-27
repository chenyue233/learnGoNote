package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("连接客户端失败,错误信息：", err)
		}
		recvStr := string(buf[:n])
		fmt.Println("收到客户端信息：", recvStr)
		conn.Write([]byte(recvStr)) //发送数据
	}
}
