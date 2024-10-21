package main

import (
	"fmt"
	"net/http"
	"voxcon/api"
	"voxcon/socket"
)

func main() {

	http.HandleFunc("/", api.HealthCheck)
	http.HandleFunc("/ws", socket.HandleConnection)

	fmt.Println("Server started on port 7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		return
	}
}
