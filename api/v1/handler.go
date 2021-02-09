package v1

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"gitlab.udevs.io/macbro/mb_corporate_service/config"
	"gitlab.udevs.io/macbro/mb_corporate_service/mb_variables"
	"gitlab.udevs.io/macbro/mb_corporate_service/pkg/logger"
	"gitlab.udevs.io/macbro/mb_corporate_service/storage"
)

type handlerV1 struct {
	log     logger.Logger
	cfg     *config.Config
	storage storage.StorageI
}

type HandlerV1Options struct {
	Log     logger.Logger
	Cfg     *config.Config
	Storage storage.StorageI
}

func New(options *HandlerV1Options) *handlerV1 {
	return &handlerV1{
		log:     options.Log,
		cfg:     options.Cfg,
		storage: options.Storage,
	}
}

func (h *handlerV1) HandleError(c *gin.Context, err error, message string) {
	if err == sql.ErrNoRows {
		h.HandleNotFoundError(c, err, message)
	} else {
		h.HandleInternalServerError(c, err, message)
	}
}

func (h *handlerV1) HandleBadRequest(c *gin.Context, err error, message string) {
	c.JSON(http.StatusNotFound, mb_variables.Error{
		Code:    http.StatusNotFound,
		Message: message,
		Reason:  err.Error(),
	})
}

func (h *handlerV1) HandleInternalServerError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusInternalServerError, mb_variables.Error{
		Code:    http.StatusInternalServerError,
		Message: message,
		Reason:  err.Error(),
	})
}

func (h *handlerV1) HandleNotFoundError(c *gin.Context, err error, message string) {
	c.JSON(http.StatusInternalServerError, mb_variables.Error{
		Code:    http.StatusInternalServerError,
		Message: message,
		Reason:  err.Error(),
	})
}
