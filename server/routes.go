package server

import (
	"net/http"
	"voxcon/socket"
)

func SetupRoutes() {
	http.HandleFunc("/", HealthCheck)
	http.HandleFunc("/game", NewGame)
	http.HandleFunc("/ws", socket.HandleConnection)
}
