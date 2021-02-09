package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
	"gitlab.udevs.io/macbro/mb_corporate_service/api"
	"gitlab.udevs.io/macbro/mb_corporate_service/config"
	"gitlab.udevs.io/macbro/mb_corporate_service/events/corporate_service"
	event "gitlab.udevs.io/macbro/mb_corporate_service/events/handler"
	"gitlab.udevs.io/macbro/mb_corporate_service/pkg/logger"
	"gitlab.udevs.io/macbro/mb_corporate_service/storage"
)

var (
	// ExchangeName ...
	ExchangeName string
	// QueueName ...
	QueueName string
	// RoutingKey - Routing Key
	RoutingKey string
)

func init() {
	flag.StringVar(&QueueName, "queue", "application.queue", "Queue name")
	flag.StringVar(&ExchangeName, "exchange_name", "application", "Queue name")
	flag.StringVar(&RoutingKey, "routing_key", "*", "Queue name")
}

func main() {
	cfg := config.Load()
	log := logger.New(cfg.Environment, "mb_corporate_service")

	defer func() {
		if err := recover(); err != nil {
			log.Error("Something went wrong", logger.Error(err.(error)))
			os.Exit(1)
		}
	}()

	psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)

	if err != nil {
		panic(err)
	}

	rabbitMQUrl := fmt.Sprintf("amqp://%s:%s@%s:%d", cfg.RabbitMQUser, cfg.RabbitMQPassword, cfg.RabbitMQHost, cfg.RabbitMQPort)

	rabbitMQConn, err := amqp.Dial(rabbitMQUrl)

	if err != nil {
		panic(err)
	}

	channel, err := rabbitMQConn.Channel()

	if err != nil {
		panic(err)
	}

	rmq, err := event.NewRMQ(context.Background(), rabbitMQConn)

	if err != nil {
		panic(err)
	}

	companyService := corporate_service.NewCompanyService(psqlConn, log, channel)

	err = rmq.NewConsumer("", "macbro", "corporate.company.create", "corporate.company.create", companyService.Create)
	if err != nil {
		log.Error("error while delete company create consumer")
		return
	}

	err = rmq.NewConsumer("", "macbro", "corporate.company.update", "corporate.company.update", companyService.Update)

	if err != nil {
		panic(err)
	}

	err = rmq.NewConsumer("", "macbro", "corporate.company.delete", "corporate.company.delete", companyService.Delete)

	if err != nil {
		log.Error("error while delete company delete consumer")
		return
	}

	go rmq.RunConsumers(context.Background())

	strg := storage.NewStoragePg(psqlConn)

	server := api.New(&api.RouterOptions{
		Log:     log,
		Cfg:     &cfg,
		Storage: strg,
	})

	_ = server.Run(cfg.HttpPort)
}
