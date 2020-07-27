package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen(Server_Type, Server_Addr)
	if err == nil {
		for {
			conn, err := listener.Accept()
			if err == nil {
				go handleConn(conn)
			}
		}
	} else {
		fmt.Println("server error ", err)
	}
	defer listener.Close()
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetReadDeadline(time.Now().Add(time.Second * 2))
		if str, err := Read(conn); err == nil {
			fmt.Println("client", conn.RemoteAddr(), str)
			Write(conn, "server got"+str)
		}
	}
}
