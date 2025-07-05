package Articles

import (
	"articles-consumer/Controller/Dto"
	"articles-consumer/Controller/Dto/Request"
	"articles-consumer/Helper"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
)

type ArticlesControllerInterface interface {
	CreateArticles(data amqp.Delivery, rabbitChannel *amqp.Channel)
}

func NewController(u Dto.Utilities) ArticlesConsumer {
	return ArticlesConsumer{u}
}

func (a ArticlesConsumer) CreateArticles(data amqp.Delivery, rabbitChannel *amqp.Channel) {
	var params Request.CreateArticle
	err := json.Unmarshal(data.Body, &params)
	if err != nil {
		responseData := Helper.SetResponsePublisher(err.Error(), nil, http.StatusBadRequest)
		err = a.RabbitMQ.ReplyRPC(data, rabbitChannel, responseData)
		if err != nil {
			log.Println(err.Error())
			return
		}

		return
	}

	err = a.Modules.Articles.CreateArticle(params)
	if err != nil {
		responseData := Helper.SetResponsePublisher(err.Error(), nil, http.StatusBadRequest)
		err = a.RabbitMQ.ReplyRPC(data, rabbitChannel, responseData)
		if err != nil {
			log.Println(err.Error())
			return
		}

		return
	}

	responseData := Helper.SetResponsePublisher("", "OK", http.StatusOK)
	err = a.RabbitMQ.ReplyRPC(data, rabbitChannel, responseData)
	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}
