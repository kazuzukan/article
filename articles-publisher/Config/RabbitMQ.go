package Config

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"os"
)

func (d connection) RabbitMQConnection() *amqp.Connection {
	return d.rabbitmq
}

func newRabbitMQConnection() *amqp.Connection {
	host := os.Getenv("HOST_RABBIT")
	if host == "" {
		host = "localhost"
	}

	conn, err := amqp.DialConfig("amqp://guest:guest@"+host+":5672/", amqp.Config{})
	if err != nil {
		log.Printf("%s: %s", err, "Failed to connect to RabbitMQ")
		panic(err)
	}

	fmt.Println("Connected to RabbitMQ")

	return conn
}
