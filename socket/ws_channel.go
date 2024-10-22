package socket

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"voxcon/constant"
	"voxcon/game"
	"voxcon/player"
)

type ChanMessage struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func handleInputChannel(ctx context.Context, in chan string, game *game.Game, player *player.Player) {
	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("Input channel closed.")
				return
			}
		case msg, ok := <-in:
			{
				if !ok {
					fmt.Println("Input channel closed.")
					return
				}

				var message *ChanMessage
				if err := json.Unmarshal([]byte(msg), &message); err != nil {
					fmt.Printf("Error unmarshalling main message: %v", err)
					break
				}

				switch message.Type {
				case constant.OfferSDP:
					{

						break
					}
				case constant.AnswerSDP:
					{

						break
					}
				}
			}
		}
	}
}

func handleOutputChannel(ctx context.Context, out chan string, conn *websocket.Conn, game *game.Game, player *player.Player) {
	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("Output channel closed.")
				return
			}
		case msg, ok := <-out:
			{
				if !ok {
					fmt.Println("Output channel closed.")
					return
				}

				if err := conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
					log.Printf("Error sending message: %v", err)
					return
				}
			}
		}
	}
}

func handleMainChannel(ctx context.Context, main chan string, cancel context.CancelFunc, conn *websocket.Conn, game *game.Game, player *player.Player) {
	for {
		select {
		case <-ctx.Done():
			{
				fmt.Println("Main channel closed.")
				return
			}
		case msg, ok := <-main:
			{
				if !ok {
					fmt.Println("Main channel closed.")
					return
				}

				var message *ChanMessage
				if err := json.Unmarshal([]byte(msg), &message); err != nil {
					fmt.Printf("Error unmarshalling main message: %v", err)
					break
				}

				switch message.Type {
				case constant.PlayerLeftWscl:
					{
						if err := conn.Close(); err != nil {
							log.Printf("Error closing connection: %v", err)
						}
						cancel()
						break
					}
				}
			}
		}
	}
}
