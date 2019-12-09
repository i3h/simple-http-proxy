package main

import (
	"net"
)

func init() {
	init_flag()
	init_log()
}

func main() {
	// listen
	l, err := net.Listen("tcp", LISTEN_ADDR+":"+LISTEN_PORT)
	check(err)
	defer l.Close()
	Log.Info("Listening on " + LISTEN_ADDR + ":" + LISTEN_PORT)

	// accept
	for {
		client, err := l.Accept()
		check(err)
		go handleConnection(client)
	}
}
