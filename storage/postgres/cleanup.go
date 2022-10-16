package postgres

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/abdukhashimov/golang-hex-architecture/storage/repo"
	"gorm.io/gorm"
)

type cleanup struct {
	db *gorm.DB
}

func NewCleanupStorage(db *gorm.DB) repo.CleanupI {
	return &cleanup{db: db}
}

func (c *cleanup) Cleanup() error {
	err := config.ConnectDB().Where("1 = 1").Delete(&models.User{}).Error
	err = config.ConnectDB().Where("1 = 1").Delete(&models.Book{}).Error

	return err
}
