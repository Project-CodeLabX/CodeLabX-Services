package ws

import (
	"codelabx/rds"
	"codelabx/rmq"
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
)

var (
	webSocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type manager struct {
	Clients map[string]client
	RmqConn *amqp.Connection
	sync.RWMutex
	Rdb *redis.Client
}

func NewManager() *manager {
	rmqConn := rmq.ConnectToRmq()
	return &manager{Clients: map[string]client{}, RmqConn: rmqConn, Rdb: rds.GetRedisClient()}
}

func (m *manager) AddClient(cl *client) {
	m.Lock()
	m.Clients[cl.Username] = *cl
	m.Unlock()
}

func (m *manager) RemoveClient(cl *client) {
	m.Lock()
	delete(m.Clients, cl.Username)
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

func (m *manager) ListenToRedis() {

	ctx := context.Background()

	for {
		for key, conn := range m.Clients {
			stdout, err := m.Rdb.Get(ctx, key).Result()
			if err == redis.Nil {
				// fmt.Println("key does not exist")
			} else if err != nil {
				log.Println("error in listenToRedis : ", err)
			} else {
				conn.WsConn.WriteMessage(1, []byte(stdout))
				m.Rdb.Del(ctx, key)
			}
		}
	}

}
