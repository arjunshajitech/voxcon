package player

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v4"
	"sync"
)

type Player struct {
	Id           string                 `json:"id"`
	GameId       string                 `json:"game_id"`
	Name         string                 `json:"name"`
	WsConn       *websocket.Conn        `json:"ws_conn"`
	SenderPeer   *webrtc.PeerConnection `json:"sender_peer"`
	ReceiverPeer *webrtc.PeerConnection `json:"receiver_peer"`
	DataChannel  *webrtc.DataChannel    `json:"data_channel"`
	InChan       chan string            `json:"in_chan"`
	OutChan      chan string            `json:"out_chan"`
	MainChan     chan string            `json:"main_chan"`
	Mu           sync.RWMutex           `json:"rwm"`
}

func NewPlayer(name string, gameId string, conn *websocket.Conn,
	inChan chan string, outChan chan string, mainChan chan string) *Player {

	return &Player{
		Id:       uuid.New().String(),
		GameId:   gameId,
		Name:     name,
		WsConn:   conn,
		MainChan: mainChan,
		InChan:   inChan,
		OutChan:  outChan,
	}
}

func (p *Player) SetSenderPeer(peer *webrtc.PeerConnection) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.SenderPeer = peer
}

func (p *Player) SetReceiverPeer(peer *webrtc.PeerConnection) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.ReceiverPeer = peer
}

func (p *Player) SetDataChannel(dc *webrtc.DataChannel) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.DataChannel = dc
}

func (p *Player) ID() string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.Id
}

func (p *Player) GetSenderPeer() *webrtc.PeerConnection {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.SenderPeer
}

func (p *Player) GetReceiverPeer() *webrtc.PeerConnection {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.ReceiverPeer
}

func (p *Player) GetDataChannel() *webrtc.DataChannel {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.DataChannel
}

func (p *Player) GetGameId() string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.GameId
}

func (p *Player) GetMainChan() chan string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.MainChan
}

func (p *Player) GetOutChan() chan string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.OutChan
}

func (p *Player) GetInChan() chan string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.InChan
}

func (p *Player) GetWsConn() *websocket.Conn {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.WsConn
}
