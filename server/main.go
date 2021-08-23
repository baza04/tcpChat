package main

import (
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	openConnections = make(map[net.Conn]bool)
	newConnection   = make(chan net.Conn)
	deadConnection  = make(chan net.Conn)
)

func main() {
	lnr, err := net.Listen("tcp", ":8080")
	logFatal(err)
	fmt.Println("listen tcp:8080")
	defer lnr.Close()

	go func() {
		for {
			conn, err := lnr.Accept()
			logFatal(err)

			openConnections[conn] = true
			newConnection <- conn
		}
	}()
	fmt.Println(<-newConnection)
}
