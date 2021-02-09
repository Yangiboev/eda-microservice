package corporate_service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.udevs.io/macbro/mb_corporate_service/mb_variables/corporate_service"

	"github.com/jmoiron/sqlx"
	"github.com/streadway/amqp"
	"gitlab.udevs.io/macbro/mb_corporate_service/mb_variables"
	"gitlab.udevs.io/macbro/mb_corporate_service/pkg/helper"
	"gitlab.udevs.io/macbro/mb_corporate_service/pkg/logger"
	"gitlab.udevs.io/macbro/mb_corporate_service/storage"
)

type companyService struct {
	storage storage.StorageI
	logger  logger.Logger
	ch      *amqp.Channel
}

func NewCompanyService(db *sqlx.DB, log logger.Logger, ch *amqp.Channel) *companyService {
	return &companyService{
		storage: storage.NewStoragePg(db),
		logger:  log,
		ch:      ch,
	}
}

func (c *companyService) Create(delivery amqp.Delivery) {
	var (
		company corporate_service.Company
	)
	fmt.Println(delivery)
	fmt.Println("create")
	err := json.Unmarshal(delivery.Body, &company)

	if err != nil {
		e := helper.HandleUnmarshallingError(c.logger, err, "error while unmarshalling", string(delivery.Body))
		bytes, _ := json.Marshal(mb_variables.Response{Error: e})
		_ = c.ch.Publish(
			"",
			delivery.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bytes,
			},
		)
		return
	}

	_, err = c.storage.Company().Create(&company)

	if err != nil {
		e := helper.HandleDBError(c.logger, err, "error while creating", string(delivery.Body))
		bytes, _ := json.Marshal(mb_variables.Response{Error: e})
		_ = c.ch.Publish(
			"",
			delivery.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bytes,
			},
		)
		return
	}
	bytes, _ := json.Marshal(mb_variables.Response{ID: company.ID, Error: mb_variables.Error{Code: http.StatusCreated}})
	_ = c.ch.Publish(
		"",
		delivery.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	)
}

func (c *companyService) Update(delivery amqp.Delivery) {
	var (
		company corporate_service.Company
	)
	fmt.Println(delivery)
	fmt.Println("update")
	err := json.Unmarshal(delivery.Body, &company)

	if err != nil {
		e := helper.HandleUnmarshallingError(c.logger, err, "error while unmarshalling", string(delivery.Body))
		bytes, _ := json.Marshal(mb_variables.Response{Error: e})
		_ = c.ch.Publish(
			"",
			delivery.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bytes,
			},
		)
		return
	}

	err = c.storage.Company().Update(&company)

	if err != nil {
		e := helper.HandleDBError(c.logger, err, "error while creating", string(delivery.Body))
		bytes, _ := json.Marshal(mb_variables.Response{Error: e})
		_ = c.ch.Publish(
			"",
			delivery.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bytes,
			},
		)
		return
	}
	bytes, _ := json.Marshal(mb_variables.Response{ID: company.ID, Error: mb_variables.Error{Code: http.StatusCreated}})
	_ = c.ch.Publish(
		"",
		delivery.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	)
}

func (c *companyService) Delete(delivery amqp.Delivery) {
	ID := delivery.UserId
	fmt.Println("delete")
	fmt.Println("delete")
	fmt.Println("delete")

	err := c.storage.Company().Delete(ID)

	if err != nil {
		e := helper.HandleDBError(c.logger, err, "error while deleting", delivery.UserId)
		bytes, _ := json.Marshal(mb_variables.Response{Error: e})
		_ = c.ch.Publish(
			"",
			delivery.ReplyTo,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        bytes,
			},
		)
		return
	}
	bytes, _ := json.Marshal(mb_variables.Response{ID: ID, Error: mb_variables.Error{Code: http.StatusAccepted}})
	_ = c.ch.Publish(
		"",
		delivery.ReplyTo,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        bytes,
		},
	)
}
