package space

import (
	"github.com/google/uuid"
	"sync"
	"voxcon/game"
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
