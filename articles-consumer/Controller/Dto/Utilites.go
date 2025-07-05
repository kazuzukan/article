package Dto

import (
	"articles-consumer/Modules"
	"articles-consumer/Service/RabbitMQ"
)

type Utilities struct {
	RabbitMQ RabbitMQ.RabbitMQInterface
	Modules  Modules.Module
}
