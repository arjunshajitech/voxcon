package internal

import (
	"github.com/google/uuid"
	"sync"
)

type Space struct {
	Id    string           `json:"id"`
	Games map[string]*Game `json:"games"`
	Mu    sync.RWMutex     `json:"rwm"`
}

func NewSpace() *Space {
	return &Space{
		Id:    uuid.New().String(),
		Games: make(map[string]*Game),
	}
}

func (s *Space) AddGame(game *Game) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	s.Games[game.Id] = game
}

func (s *Space) RemoveGame(game *Game) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	delete(s.Games, game.Id)
}

func (s *Space) GetGame(id string) *Game {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return s.Games[id]
}

func (s *Space) GetAllGames() []*Game {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	var games []*Game
	for _, game := range s.Games {
		games = append(games, game)
	}
	return games
}
