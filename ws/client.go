package ws

import (
	"codelabx/rmq"
	"log"

	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
)

type client struct {
	Username string
	WsConn   *websocket.Conn
	Manager  *manager
	RmqCh    *amqp.Channel
}

func NewClient(username string, wsConn *websocket.Conn, manager *manager) client {
	rmqCh := rmq.CreateChannel(manager.RmqConn)
	return client{Username: username, WsConn: wsConn, Manager: manager, RmqCh: rmqCh}
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
