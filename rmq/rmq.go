package rmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

const (
	url string = "amqps://abhi:Deadshot1060@b-195dfc46-2db6-4582-b92b-ff6bc1a3b4fd.mq.ap-south-1.amazonaws.com:5671/codelabx"
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
