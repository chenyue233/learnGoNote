package api

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// ListenTCP(net string, laddr *TCPAddr) (l *TCPListener, err os.Error)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	// func (l *TCPListener) Accept() (c Conn, err os.Error)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	daytime := time.Now().String()
	// don't care about return value
	conn.Write([]byte(daytime))

	// we're finished with this client
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
