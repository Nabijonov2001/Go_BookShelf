package v1

import (
	"net/http"

	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type handlerV1 struct {
	log     *zap.Logger
	cfg     *config.Config
	service service.ServiceI
}

type HandlerOptions struct {
	Cfg     *config.Config
	Log     *zap.Logger
	Service service.ServiceI
}

func NewHandler(opts *HandlerOptions) *handlerV1 {
	return &handlerV1{
		cfg:     opts.Cfg,
		log:     opts.Log,
		service: opts.Service,
	}
}

func (h *handlerV1) handleInternal(c *gin.Context, msg string, err error) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseError{
			Code:    "INTERNAL_SERVER_ERR",
			Message: msg,
			Error:   err.Error(),
		})
		return true
	}
	return false
}

func (h *handlerV1) handleBadRequest(c *gin.Context, msg string, err error) bool {
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseError{
			Code:    "BAD_REQUEST",
			Message: msg,
		})
		return true
	}
	return false
}

type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Error   string `json:"error"`
}

type ResponseCreated struct {
	IsOk    bool        `json:"isOk"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
