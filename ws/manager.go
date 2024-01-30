package ws

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	webSocketUpgrader websocket.Upgrader
)

type manager struct{}

func NewManager() *manager {
	return &manager{}
}

func (m *manager) ServeWs(w http.ResponseWriter, r *http.Request) {
	_, err := webSocketUpgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println("err during upgrade")
		return
	}
	fmt.Println("client connected")
}
