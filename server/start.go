package server

import (
	"Nvwa/server/nvwa_http"
	"Nvwa/server/nvwa_socket"
)

type Server interface {
	StartServer(address string)
}

func StartSocketServer(connectionType string, address string) {
	nvwa_socket.InitSocketHandler()
	var server Server
	switch connectionType {
	case "ws":
		server = &nvwa_socket.WebSocketServer{}
	default:
		server = &nvwa_socket.WebSocketServer{}
	}
	server.StartServer(address)
}

func StartHttpServer(address string) {
	server := &nvwa_http.HttpServer{}
	server.StartServer(address)
}
