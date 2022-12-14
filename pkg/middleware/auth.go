package middleware

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/pkg/helper"
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {

	if len(c.Request.Header["Key"]) == 0 || len(c.Request.Header["Sign"]) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"code":    "INTERNAL_SERVER_ERR",
			"message": "access denied",
		})
		return
	}

	var user models.User
	key := c.Request.Header["Key"][0]
	sign := c.Request.Header["Sign"][0]

	result := config.ConnectDB().Where("key=?", key).First(&user)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "UNAUTHORIZED",
			"message": "access denied",
		})
		return
	}

	bodyByte, err := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyByte))
	if err != nil {
		bodyByte = []byte("")
	}

	hash := helper.CreateHash(c.Request.Method + config.GetEnv("BASE_URL") + c.Request.RequestURI + string(bodyByte) + user.Secret)

	if sign != hash {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"code":    "UNAUTHORIZED",
			"message": "access denied",
		})
		return
	}

	c.Set("user", user)
	c.Next()

}
