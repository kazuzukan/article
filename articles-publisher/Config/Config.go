package Config

import (
	"database/sql"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ConnectionInterface interface {
	PostgreSQLConnection() *sql.DB
	RabbitMQConnection() *amqp.Connection
}

type connection struct {
	postgresql *sql.DB
	rabbitmq   *amqp.Connection
}

func BuildConnection() connection {
	return connection{
		postgresql: newConnectionPostgreSQL(),
		rabbitmq:   newRabbitMQConnection(),
	}
}
