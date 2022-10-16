package v1

import (
	"net/http"

	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/gin-gonic/gin"
)

func (h *handlerV1) GetAllBooks(c *gin.Context) {
	var response ResponseCreated

	res, err := h.service.Book().GetAllBooks()
	if h.handleBadRequest(c, "failed to get all books", err) {
		return
	}

	response.IsOk = true
	response.Message = "ok"
	response.Data = res

	c.JSON(http.StatusOK, response)
}

func (h *handlerV1) CreateBook(c *gin.Context) {
	var (
		payload  models.BookCreate
		response ResponseCreated
	)

	err := c.ShouldBindJSON(&payload)
	if h.handleBadRequest(c, "failed to convert json to struct", err) {
		return
	}

	res, err := h.service.Book().CreateBook(payload)
	if h.handleBadRequest(c, "failed to create book", err) {
		return
	}

	response.IsOk = true
	response.Data = res
	response.Message = "ok"

	c.JSON(http.StatusOK, response)
}

func (h *handlerV1) UpdateBook(c *gin.Context) {
	var (
		payload  models.BookUpdate
		response ResponseCreated
	)

	err := c.ShouldBindJSON(&payload)
	if h.handleBadRequest(c, "failed to convert json to struct", err) {
		return
	}

	res, err := h.service.Book().UpdateBook(c.Param("id"), payload)
	if h.handleBadRequest(c, "failed to update a book", err) {
		return
	}

	response.IsOk = true
	response.Message = "ok"
	response.Data = res

	c.JSON(http.StatusOK, response)
}

func (h *handlerV1) DeleteBook(c *gin.Context) {
	var response ResponseCreated

	err := h.service.Book().DeleteBook(c.Param("id"))
	if h.handleBadRequest(c, "failed to delete a book", err) {
		return
	}

	response.IsOk = true
	response.Message = "ok"
	response.Data = "Successfully deleted"

	c.JSON(http.StatusOK, response)
}
