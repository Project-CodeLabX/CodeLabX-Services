package rmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	url string = os.Getenv("CLX_MQ")
)

func ConnectToRmq() *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Println("error in rmq connection")
	}
	log.Println("Connected to rmq...")
	return conn
}

func CreateChannel(conn *amqp.Connection) *amqp.Channel {
	ch, err := conn.Channel()
	if err != nil {
		log.Println("error in rmq channel creation")
	}
	return ch
}
