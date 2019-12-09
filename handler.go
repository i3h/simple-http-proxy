package main

import (
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func handleHTTPConnection(client net.Conn) {
	b := make([]byte, 2048)
	n, _ := client.Read(b)
	words := strings.Fields(string(b))
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
		server.Write(b[:n])
	}

	go io.Copy(server, client)
	io.Copy(client, server)
}

func handleSOCKS5Connection(client net.Conn) {
	b := make([]byte, 2048)
	n, _ := client.Read(b)
	if b[0] == 0x05 {
		client.Write([]byte{0x05, 0x00})
		n, _ = client.Read(b)
		var host, port string
		switch b[3] {
		case 0x01:
			host = net.IPv4(b[4], b[5], b[6], b[7]).String()
		case 0x03:
			host = string(b[5 : n-2])
		case 0x04:
			host = net.IP{b[4], b[5], b[6], b[7],
				b[8], b[9], b[10], b[11],
				b[12], b[13], b[14], b[15],
				b[16], b[17], b[18], b[19]}.String()
		}
		port = strconv.Itoa(int(b[n-2])<<8 | int(b[n-1]))

		server, err := net.Dial("tcp", net.JoinHostPort(host, port))
		if err != nil {
			Log.Warn(err)
			return
		}
		defer server.Close()
		client.Write([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

		go io.Copy(server, client)
		io.Copy(client, server)
	}
}
