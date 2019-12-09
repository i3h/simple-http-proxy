package main

import (
	"io"
	"net"
	"net/url"
	"strings"
	"time"
)

func handleConnection(client net.Conn) {
	buf := make([]byte, 2048)
	reqLen, _ := client.Read(buf)
	words := strings.Fields(string(buf))
	if len(words) < 3 {
		client.Close()
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

	d := net.Dialer{Timeout: 5 * time.Second}
	server, err := d.Dial("tcp", address)

	if err != nil {
		Log.Warn("ERR", "    ", address)
		client.Close()
		return
	}

	if method == "CONNECT" {
		client.Write([]byte("HTTP/1.1 200 Connection established\r\n\r\n"))
	} else {
		server.Write(buf[:reqLen])
	}

	go io.Copy(server, client)
	io.Copy(client, server)
}
