package RabbitMQ

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQInterface interface {
	PublishMessageRPC(key, exchange string, data interface{}) (msg []byte, err error)
}

type RabbitMQService struct {
	RabbitCon *amqp.Connection
}

func GetRabbitMQService(con *amqp.Connection) RabbitMQService {
	return RabbitMQService{RabbitCon: con}
}
