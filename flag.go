package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"strconv"
)

type ConfigData struct {
	Proxy_Type  string `json:"proxy_type"`
	Listen_Addr string `json:"listen_addr"`
	Listen_Port int    `json:"listen_port"`
}

var Config ConfigData

var PROXY_TYPE string
var LISTEN_ADDR string
var LISTEN_PORT string

func init_flag() {
	// Read config.json
	file := flag.String("c", "config_test.json", "Location of the config file.")
	flag.Parse()
	jsonData, err := ioutil.ReadFile(*file)
	check(err)
	err = json.Unmarshal(jsonData, &Config)
	check(err)
	// Rename
	PROXY_TYPE = Config.Proxy_Type
	LISTEN_ADDR = Config.Listen_Addr
	LISTEN_PORT = strconv.Itoa(Config.Listen_Port)
}
