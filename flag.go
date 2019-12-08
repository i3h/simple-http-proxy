package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"strconv"
)

type ConfigData struct {
	Server string `json:"server"`
	Port   int    `json:"port"`
	Type   string `json:"type"`
}

var Config ConfigData

var CONN_HOST string
var CONN_PORT string
var CONN_TYPE string

func init_flag() {
	// Read config.json
	file := flag.String("c", "config.json", "Location of the config file.")
	flag.Parse()
	jsonData, err := ioutil.ReadFile(*file)
	check(err)
	err = json.Unmarshal(jsonData, &Config)
	check(err)
	// Rename
	CONN_HOST = Config.Server
	CONN_PORT = strconv.Itoa(Config.Port)
	CONN_TYPE = Config.Type
}
