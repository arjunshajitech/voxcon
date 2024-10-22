package space

import (
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
	"voxcon/constant"
	"voxcon/game"
	"voxcon/server"
	"voxcon/socket"
)

type Space struct {
	Id    string                `json:"id"`
	Games map[string]*game.Game `json:"games"`
	Mu    sync.RWMutex          `json:"rwm"`
}

func NewSpace() *Space {
	return &Space{
		Id:    uuid.New().String(),
		Games: make(map[string]*game.Game),
	}
}

func (s *Space) Start() {

	// TODO: Need to fix
	g1 := game.NewGame(constant.DefaultGameID)
	s.AddGame(g1)
	go s.Log()

	http.HandleFunc("/", server.HealthCheck)
	http.HandleFunc("/ws", s.handleConnection)

	fmt.Println("Server started on port 7777")
	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		return
	}
}

func (s *Space) handleConnection(w http.ResponseWriter, r *http.Request) {
	socket.HandleConnection(w, r, s)
}

func (s *Space) AddGame(game *game.Game) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Games[game.Id] = game
}

func (s *Space) RemoveGame(game *game.Game) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Games, game.Id)
}

func (s *Space) GetGame(id string) *game.Game {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return s.Games[id]
}

func (s *Space) GetAllGames() []*game.Game {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	var games []*game.Game
	for _, g := range s.Games {
		games = append(games, g)
	}
	return games
}

func (s *Space) Log() {
	ticker := time.NewTicker(5000 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			{
				s.Mu.RLock()
				games := s.GetAllGames()
				s.Mu.RUnlock()
				for _, g := range games {
					fmt.Printf("GameID (%s) PlayersCount (%v)\n", g.Id, len(g.GetPlayers()))
				}
			}
		}
	}

}
