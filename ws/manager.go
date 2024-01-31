package ws

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type manager struct {
	Clients map[string]client
	sync.RWMutex
}

func NewManager() *manager {
	return &manager{Clients: map[string]client{}}
}

func (m *manager) AddClient(cl *client) {
	m.Lock()
	m.Clients[cl.Username] = *cl
	m.Unlock()
}

func (m *manager) ServeWs(w http.ResponseWriter, r *http.Request) {
	wsConn, err := webSocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("err during upgrade")
		return
	}
	fmt.Println("client connected")
	username := r.Header.Get("username")

	cl := NewClient(username, wsConn, m)

	m.AddClient(&cl)

	go cl.ListenToClient()
}
