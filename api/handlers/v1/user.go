package v1

import (
	"net/http"

	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) GetOneUser(c *gin.Context) {
	var response ResponseCreated

	user, _ := c.Get("user")

	response.IsOk = true
	response.Data = user
	response.Message = "ok"

	c.JSON(http.StatusOK, response)
}

func (h *handlerV1) CreateUser(c *gin.Context) {
	var (
		payload  models.User
		response ResponseCreated
	)

	err := c.ShouldBindJSON(&payload)
	if h.handleBadRequest(c, "failed to convert json to struct", err) {
		return
	}

	res, err := h.service.User().CreateUser(payload)
	if h.handleInternal(c, "failed to signup user", err) {
		return
	}

	response.IsOk = true
	response.Data = res
	response.Message = "ok"

	c.JSON(http.StatusOK, response)
}
