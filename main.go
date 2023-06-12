package main

import (
	"Nvwa/game/spirit/agent"
	"Nvwa/logger"
	"Nvwa/server"
)

func main() {
	logger.InitNvwaLogger("debug")
	agent.Init()
	server.StartHttpServer(":8888")
}
