package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// make a buffer to hold incoming data
	buf := make([]byte, 1024)
	// read the incoming connection into the buffer
	//reqLen, err := conn.Read(buf)
	_, err := conn.Read(buf)
	//fmt.Println(reqLen)
	check(err)
	fmt.Println(string(buf))

	// send a response back
	conn.Write([]byte("message received"))
	// close the connection when you are done with it
	conn.Close()
}
