package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handlerV1) Get(c *gin.Context) {
	id := c.Param("id")

	_, err := uuid.Parse(id)

	if err != nil {
		h.HandleBadRequest(c, err, "Id format should be uuid")
		return
	}

	resp, err := h.storage.Company().Get(c.Param("id"))

	if err != nil {
		h.HandleError(c, err, "Company not found")
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handlerV1) GetAll(c *gin.Context) {
	resp, count, err := h.storage.Company().GetAll(0, 0, c.Query("name"))

	if err != nil {
		h.HandleInternalServerError(c, err, "Something went wrong")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"companies": resp,
		"count":     count,
	})
}
