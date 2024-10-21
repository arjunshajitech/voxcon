package game

import (
	"github.com/google/uuid"
	"sync"
	"time"
	"voxcon/player"
)

type Game struct {
	Id        string                    `json:"id"`
	HostId    string                    `json:"hostId"`
	Players   map[string]*player.Player `json:"players"`
	CreatedAt time.Time                 `json:"createdAt"`
	Mu        sync.RWMutex              `json:"rwm"`
}

func NewGame(id string) *Game {
	return &Game{
		Id:        id,
		HostId:    uuid.New().String(),
		Players:   make(map[string]*player.Player),
		CreatedAt: time.Now(),
	}
}

func (g *Game) ID() string {
	g.Mu.RLock()
	defer g.Mu.RUnlock()
	return g.Id
}

func (g *Game) GetPlayer(id string) *player.Player {
	g.Mu.RLock()
	defer g.Mu.RUnlock()
	return g.Players[id]
}

func (g *Game) SetPlayer(id string, player *player.Player) {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	g.Players[id] = player
}
