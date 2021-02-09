package event

import (
	"github.com/streadway/amqp"
)

type Consumer struct { //Subscribe
	channel      *amqp.Channel
	consumerName string
	exchangeName string
	routingKey   string
	queueName    string
	middlewares  []func(amqp.Delivery) error
	messages     <-chan amqp.Delivery
	handler      func(amqp.Delivery)
}

func (c *Consumer) AppendMiddleware(f func(delivery amqp.Delivery) error) {
	c.middlewares = append(c.middlewares, f)
}
