package main

import (
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
		conn.Close()
		return
	}
	method := words[0]
	rawUrl := words[1]
	var hostname string
	var address string
	if method == "CONNECT" {
		hostname = strings.Split(rawUrl, ":")[0]
		address = hostname + ":443"
		Log.Info(method[0:3], "    ", address)
	} else {
		u, _ := url.Parse(rawUrl)
		hostname = u.Host
		address = hostname + ":80"
		Log.Info(method[0:3], "    ", address)
	}

	dialer, err := net.Dial("tcp", address)
	if err != nil {
		Log.Info("dail failed")
		conn.Close()
		return
	}

	if method == "CONNECT" {
		conn.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	} else {
		dialer.Write(buf[:reqLen])
	}

	go io.Copy(dialer, conn)
	io.Copy(conn, dialer)
}
