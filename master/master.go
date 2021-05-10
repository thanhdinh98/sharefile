package master

import (
	"fmt"
	"log"
	"net"
)

var (
	Address = "localhost:1234"
)

func init() {
	Addr, err := net.ResolveTCPAddr("tcp", Address)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Print(conn)
	}
}
