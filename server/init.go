package server

import (
	"net"
	"time"

	"github.com/nbanitama/chat/util"
)

// Start represents to run a server
func Start() {
	service := ":7777"
	tcpAddress, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		util.Debug("got %v", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddress)
	if err != nil {
		util.Debug("got %v", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			util.Debug("got %v", err)
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	util.Debug("handleClient: get new Client")
	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}
