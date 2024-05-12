package ws

import (
	"codelabx/rmq"
	"context"
	"encoding/json"
	"log"
	"time"

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
		c.RmqCh.Close()
		c.Manager.RemoveClient(c)
	}()
	log.Println("started Listening to client...")

	for {
		msgType, msg, err := c.WsConn.ReadMessage()

		if err != nil {
			log.Println("Error in reading message : ", err)
			break
		}

		log.Println("From : " + c.Username + ", Message type: " + string(msgType) + ", this is msg : " + string(msg))

		var userEvent rmq.UserEvent
		er := json.Unmarshal(msg, &userEvent)
		userEvent.UserName = c.Username
		if er != nil {
			log.Println("json Unmarshal failed...")
			break
		}

		context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		c.SendUserEventWithContext(context, &userEvent)
	}
	log.Println("stopped Listening to client...")
}

func (c *client) SendUserEventWithContext(ctx context.Context, userEvent *rmq.UserEvent) error {
	event, err := json.Marshal(*userEvent)
	if err != nil {
		panic("error is send Message Marshaling...")
	}
	msg := amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Transient,
		Body:         event,
	}
	if userEvent.Language == "python" {
		return c.RmqCh.PublishWithContext(ctx, "codelabx", "py_events", true, false, msg)
	}
	return c.RmqCh.PublishWithContext(ctx, "codelabx", "java_events", true, false, msg)
}
