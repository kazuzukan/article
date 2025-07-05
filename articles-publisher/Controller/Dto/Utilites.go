package Dto

import (
	"articles-publisher/Service/Articles"
	"articles-publisher/Service/RabbitMQ"
)

type Utilities struct {
	RabbitMQ        RabbitMQ.RabbitMQInterface
	ArticlesService Articles.ArticleServiceInterface
}
