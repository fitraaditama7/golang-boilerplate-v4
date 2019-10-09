package main

import (
	"golang-websocket/api"
	"golang-websocket/config"
)

func main() {
	config.Init()
	api.Run()
}
