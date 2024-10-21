package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"voxcon/constant"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity
	},
}

func Start() {

}

func HandleConnection(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Error during connection upgrade:", err)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	inChan := make(chan string, 10)
	outChan := make(chan string, 10)
	mainChan := make(chan string, 10)

	go handleInputChannel(ctx, inChan)
	go handleOutputChannel(ctx, outChan, conn)
	go handleMainChannel(ctx, mainChan, cancel, conn)

	defer func() {
		wsclMess, _ := json.Marshal(&MainChanMessage{
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
