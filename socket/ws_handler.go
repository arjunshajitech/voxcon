package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"voxcon/constant"
	"voxcon/player"
	"voxcon/space"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity
	},
}

func HandleConnection(w http.ResponseWriter, r *http.Request, space *space.Space) {

	// TODO: Validate auth token during handshake then continue
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	inChan := make(chan string, 10)
	outChan := make(chan string, 10)
	mainChan := make(chan string, 10)

	g := space.GetGame(constant.DefaultGameID)
	p := player.NewPlayer("name", g.ID(), conn, inChan, outChan, mainChan)
	g.SetPlayer(p)

	go handleInputChannel(ctx, inChan, g, p)
	go handleOutputChannel(ctx, outChan, conn, g, p)
	go handleMainChannel(ctx, mainChan, cancel, conn, g, p)

	defer func() {
		wsclMess, _ := json.Marshal(&ChanMessage{
			Type: constant.PlayerLeftWscl,
			Data: nil,
		})
		mainChan <- string(wsclMess)
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		inChan <- string(msg)
	}
}
