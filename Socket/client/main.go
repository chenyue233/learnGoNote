package main

import (
	"bytes"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "172.0.0.1:8085")
	if err != nil {
		fmt.Println("Client create conn err:", err)
	}
	defer conn.Close()
	Write(conn, "aaaaa")
	if str, err := Read(conn); err == nil {
		fmt.Println(str)
	}
}

func Write(conn net.Conn, context string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(context)
	buffer.WriteByte('\t')
	return conn.Write(buffer.Bytes())
}

func Read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		if _, err := conn.Read(readBytes); err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == '\t' {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}
