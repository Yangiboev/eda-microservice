package helper

import (
	"database/sql"
	"gitlab.udevs.io/macbro/mb_corporate_service/mb_variables"
	"gitlab.udevs.io/macbro/mb_corporate_service/pkg/logger"
	"net/http"
)

func HandleDBError(log logger.Logger, err error, message string, req interface{}) (e mb_variables.Error) {
	if err == sql.ErrNoRows {
		log.Error(message+", Not Found", logger.Error(err), logger.Any("req", req))
		return mb_variables.Error{
			Code:    http.StatusNotFound,
			Message: message,
			Reason:  err.Error(),
		}
	} else if err != nil {
		log.Error(message, logger.Error(err), logger.Any("req", req))
		return mb_variables.Error{
			Code:    http.StatusInternalServerError,
			Message: message,
			Reason:  err.Error(),
		}
	}
	return
}
func HandleUnmarshallingError(log logger.Logger, err error, message string, req interface{}) mb_variables.Error {
	log.Error(message+", Bad Request", logger.Error(err), logger.Any("req", req))
	return mb_variables.Error{
		Code:    http.StatusBadRequest,
		Message: message,
		Reason:  err.Error(),
	}
}

//func HandlePublish()
