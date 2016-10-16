package cg

import (
    "encoding/json"
    "errors"
    "sync"
    "ipc"
)

var _ ipc.Server = &CenterServer{}

type Message struct {
    From string "from"
    To string "to"
    Content string "content"
}

type CenterServer struct {
    servers map[string] ipc.Server
    players []*Player
    rooms []*Room
    mutex sync.RWMutex
}

func NewCenterServer() *CenterServer {
    servers := make(map[string] ipc.Server)
    players := make([]*Player,0)
    return &CenterServer{servers:servers,players:players}
}

func (server *CenterServer)addPlayer(params string)error {
    player := NewPlayer()
    err := json.Unmarshal([]byte(params),&player)
    if err != nil {
        return err
    }
    server.mutex.Lock()
    defer server.mutex.Unlock()
    server.players = append(server.players,player)
    return nil
}

func (server *CenterServer)removePlayer(params string) error {
    server.mutex.Lock()
    defer server.mutex.Unlock()
    for i,v := range server.players {
        if v.Name == params {
            if len(server.players) == 1 {
                server.players = make([]*Player,0)
            } else if i == len(server.players) - 1 {
                server.players = server.players[:i - 1]
            } else if i == 0 {
                server.players = server.players[1:]
            } else {
                server.players = append(server.players[:i-1],server.players[:i + 1]...)
            }
        return nil
        }
    }
    return errors.New("Player not found")
}


