package event

import (
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"gitlab.udevs.io/macbro/mb_corporate_service/pkg/logger"
)

type RMQ struct {
	context        context.Context
	conn           *amqp.Connection
	consumers      []*Consumer
	consumerErrors chan error
	logger         logger.Logger
}

func NewRMQ(context context.Context, conn *amqp.Connection) (*RMQ, error) {
	return &RMQ{
		context:        context,
		conn:           conn,
		consumers:      []*Consumer{},
		consumerErrors: make(chan error),
	}, nil
}

func (rmq *RMQ) NewConsumer(consumerName, exchangeName, routingKey, queueName string, handler func(amqp.Delivery)) error {
	ch, err := rmq.conn.Channel()

	if err != nil {
		return err
	}

	err = declareExchange(ch, exchangeName)

	if err != nil {
		return err
	}

	queue, err := declareQueue(ch, queueName)

	if err != nil {
		return err
	}

	err = queueBindToExchange(ch, exchangeName, routingKey, queue.Name)

	if err != nil {
		return err
	}

	rmq.consumers = append(rmq.consumers, &Consumer{
		channel:      ch,
		consumerName: consumerName,
		exchangeName: exchangeName,
		routingKey:   routingKey,
		queueName:    queueName,
		middlewares:  make([]func(amqp.Delivery) error, 0),
		handler:      handler,
	})

	return nil
}

func (rmq *RMQ) RunConsumers(ctx context.Context) {
	fmt.Println("starting...")
	//info about consumers
	var err error
	for _, consumer := range rmq.consumers {
		c := consumer
		go func() {
			c.messages, err = c.channel.Consume(
				c.queueName,
				c.consumerName,
				true,
				false,
				false,
				true,
				nil,
			)

			if err != nil {
				fmt.Println(err)
				rmq.sendError(err)
			}

			for {
				select {
				case msg := <-c.messages:
					for _, middleware := range c.middlewares {
						err := middleware(msg)

						if err != nil {
							rmq.sendError(err)
						}
					}
					c.handler(msg)
				case <-rmq.context.Done(): //every
					{
						err = c.channel.Cancel("", true)
						if err != nil {
							rmq.sendError(err)
						}
						return
					}
				}
			}
		}()
	}
	go rmq.ReceiveError()
}

func (rmq *RMQ) sendError(err error) {
	rmq.consumerErrors <- err
}

func (rmq *RMQ) ReceiveError() {
	for err := range rmq.consumerErrors {
		rmq.logger.Error("error while consuming", logger.Error(err))
	}
}
