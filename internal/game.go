package internal

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

type Game struct {
	Id        string             `json:"id"`
	Name      string             `json:"name"`
	HostId    string             `json:"hostId"`
	Players   map[string]*Player `json:"players"`
	CreatedAt time.Time          `json:"createdAt"`
	Mu        sync.RWMutex       `json:"rwm"`
}

func NewGame(id, name string) *Game {
	return &Game{
		Id:        id,
		Name:      name,
		HostId:    uuid.New().String(),
		Players:   make(map[string]*Player),
		CreatedAt: time.Now(),
	}
}

func (g *Game) ID() string {
	g.Mu.RLock()
	defer g.Mu.RUnlock()
	return g.Id
}

func (g *Game) GetPlayer(id string) *Player {
	g.Mu.RLock()
	defer g.Mu.RUnlock()
	return g.Players[id]
}

func (g *Game) GetName() string {
	g.Mu.RLock()
	defer g.Mu.RUnlock()
	return g.Name
}

func (g *Game) SetPlayer(id string, player *Player) {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	g.Players[id] = player
}
