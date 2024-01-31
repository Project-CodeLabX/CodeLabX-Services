package ws

import (
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	Username string
	WsConn   *websocket.Conn
	Manager  *manager
}

func NewClient(username string, wsConn *websocket.Conn, manager *manager) client {
	return client{Username: username, WsConn: wsConn, Manager: manager}
}

func (c *client) ListenToClient() {
	defer func() {
		c.WsConn.Close()
	}()
	log.Println("started Listening to client...")
	for {
		msgType, msg, err := c.WsConn.ReadMessage()

		if err != nil {
			log.Println("Error in reading message : ", err)
			break
		}

		log.Println("From : " + c.Username + ", Message type: " + string(msgType) + ", this is msg : " + string(msg))
	}

	log.Println("stopped Listening to client...")
}
