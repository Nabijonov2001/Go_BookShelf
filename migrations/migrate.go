package main

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
)

func main() {
	config.ConnectDB().AutoMigrate(&models.User{})
}
