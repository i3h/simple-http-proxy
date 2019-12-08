package main

import (
	"fmt"
	"net"
)

func init() {
	init_flag()
}

func main() {
	server()
}

func server() {
	// listen for incoming connections
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)

	// close the listener when the application closes
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		// listen for an incoming connection
		conn, err := l.Accept()
		check(err)

		// handle connections in a new goroutine
		go handleConnection(conn)
	}
}
