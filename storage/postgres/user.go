package postgres

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/abdukhashimov/golang-hex-architecture/storage/repo"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB 
}

func NewUserStorage(db *gorm.DB) repo.UserI {
	return &users{db: db}
}

func (u *users) CreateUser(payload models.User) (models.User, error) {
	result:=config.ConnectDB().Create(&payload)
	
	return payload, result.Error
}




