package main

import (
	"voxcon/server"
	"voxcon/space"
)

func main() {
	s := space.NewSpace()
	server.Start(s)
}
