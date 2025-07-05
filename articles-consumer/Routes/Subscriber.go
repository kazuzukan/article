package Routes

import (
	Controller "articles-consumer/Controller"
	"articles-consumer/Routes/ConfigSubscriber"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type RabbitConsumerInterface interface {
	Subscriber(config ConfigSubscriber.ConfigSubs, fn func(data amqp.Delivery, rabbitChannel *amqp.Channel))
	StartSubscriber()
}

type RabbitSubscriber struct {
	Channel    *amqp.Connection
	Controller Controller.ControllerConsumer
	Config     ConfigSubscriber.SubsConfig
}

func NewRabbitMQSubscriber(rabbitCon *amqp.Connection, controller Controller.ControllerConsumer) RabbitConsumerInterface {
	return RabbitSubscriber{
		Channel:    rabbitCon,
		Controller: controller,
		Config: ConfigSubscriber.SubsConfig{
			Articles: ConfigSubscriber.ConfigSubs{},
		},
	}
}

func (r RabbitSubscriber) Subscriber(config ConfigSubscriber.ConfigSubs, fn func(data amqp.Delivery, rabbitChannel *amqp.Channel)) {
	ch, err := r.Channel.Channel()
	if err != nil {
		return
	}
	defer ch.Close()

	// exchange
	err = ch.ExchangeDeclare(
		config.ExchangeName, // name
		config.ExchangeType, // type
		true,                // durable
		false,               // auto-deleted
		false,               // internal
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		return
	}

	// queue
	q, err := ch.QueueDeclare(
		config.QueueName, // name
		false,            // durable
		false,            // delete when unused
		true,             // exclusive
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		return
	}

	// binding key
	err = ch.QueueBind(
		q.Name,              // queue name
		config.BindingKey,   // routing key
		config.ExchangeName, // exchange
		false,
		nil)

	// consumer from publisher
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		return
	}

	for d := range msgs {
		fn(d, ch)
	}

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")

}

func (r RabbitSubscriber) StartSubscriber() {
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	go r.Subscriber(r.Config.Articles.CreateArticlesConfig(), r.Controller.Articles.CreateArticles)
}
