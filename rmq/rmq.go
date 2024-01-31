package rmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	url string = "amqp://Abhi1060:Abhi1060@localhost:5672/codelabx"
)

func ConnectToRmq() *amqp.Connection {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Println("error in rmq connection")
	}
	log.Println("Connected to rmq...")
	return conn
}
