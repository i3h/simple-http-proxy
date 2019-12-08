package main

import (
	"fmt"
	"net"
)

func init() {
	init_flag()
}

func main() {
	// listen
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	check(err)
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	for {
		conn, err := l.Accept()
		check(err)
		go handleConnection(conn)
	}
}
