package main

import (
	"fmt"
	"io"
	"net"
	"net/url"
	"strings"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 2048)
	reqLen, err := conn.Read(buf)
	check(err)
	words := strings.Fields(string(buf))
	if len(words) < 3 {
		fmt.Println("length: ", len(words))
		fmt.Println(string(buf))
		//conn.Close()
		return
	}
	method := words[0]
	rawUrl := words[1]
	var hostname string
	var address string
	if method == "CONNECT" {
		//fmt.Println("method: ", method)
		//fmt.Println("host: ", strings.Split(rawUrl, ":")[0])
		hostname = strings.Split(rawUrl, ":")[0]
		address = hostname + ":443"
	} else {
		u, _ := url.Parse(rawUrl)
		//fmt.Println("method: ", method)
		//fmt.Println("host: ", u.Host)
		hostname = u.Host
		address = hostname + ":80"
	}

	dialer, err := net.Dial("tcp", address)
	check(err)

	if method == "CONNECT" {
		conn.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	} else {
		dialer.Write(buf[:reqLen])
	}

	go io.Copy(dialer, conn)
	io.Copy(conn, dialer)
}
