package RabbitMQ

import (
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (r RabbitMQService) PublishMessageRPC(key, exchange string, data interface{}) (msg []byte, err error) {
	ch, err := r.RabbitCon.Channel()
	if err != nil {
		return nil, err
	}
	defer ch.Close()

	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	corrId := uuid.NewV4().String()
	// Config Queue
	queue, err := ch.QueueDeclare(
		queueName(key),
		false,
		true,
		true,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	// consume message
	messages, err := ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	err = ch.Publish(exchange, key, false, false, amqp.Publishing{
		ContentType:   "application/json",
		Body:          dataByte,
		ReplyTo:       queue.Name,
		CorrelationId: corrId,
	})
	if err != nil {
		return nil, err
	}
	for msg := range messages {
		if corrId == msg.CorrelationId {
			return msg.Body, nil
		}
	}

	return
}

func (r RabbitMQService) ReplyRPC(delivery amqp.Delivery, rabbitChannel *amqp.Channel, data interface{}) (err error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = rabbitChannel.PublishWithContext(ctx,
		"",               // exchange
		delivery.ReplyTo, // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType:   "application/json",
			CorrelationId: delivery.CorrelationId,
			Body:          dataByte,
		})
	if err != nil {
		delivery.Ack(false)
		return
	}

	delivery.Ack(false)
	return
}

func queueName(prefix string) string {
	uuid := uuid.NewV4().String()
	return fmt.Sprintf("%s-%s", string(prefix), uuid)
}
