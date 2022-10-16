package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handlerV1) Cleanup(c *gin.Context) {
	var response ResponseCreated
	err := h.service.Cleanup().Cleanup()
	if h.handleInternal(c, "failed to clean up test data", err) {
		return
	}

	response.IsOk = true
	response.Message = "test data successfully cleaned up"

	c.JSON(http.StatusOK, response)
}
