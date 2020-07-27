package main

import (
	"bytes"
	"net"
)

const (
	Server_Type = "tcp"
	Server_Addr = "172.0.0.1:8085"
	Delimiter   = '\t'
)

func Write(conn net.Conn, context string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(context)
	buffer.WriteByte(Delimiter)
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
		if readByte == Delimiter {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}
