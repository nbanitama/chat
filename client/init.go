package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/nbanitama/chat/util"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port ", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		util.Debug("got %v", err)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		util.Debug("got %v", err)
	}
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		util.Debug("got %v", err)
	}
	result, err := ioutil.ReadAll(conn)
	if err != nil {
		util.Debug("got %v", err)
	}
	util.Debug(string(result))
	os.Exit(0)
}
