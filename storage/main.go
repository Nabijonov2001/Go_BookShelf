package storage

import (
	"github.com/abdukhashimov/golang-hex-architecture/storage/postgres"
	"github.com/abdukhashimov/golang-hex-architecture/storage/repo"
	"gorm.io/gorm"
)

type storage struct {
	userStorage repo.UserI
	bookStorage repo.BookI
}

type StorageI interface {
	User() repo.UserI
	Book() repo.BookI
}

func NewStorage(db *gorm.DB) StorageI {
	return &storage{
		userStorage: postgres.NewUserStorage(db),
		bookStorage: postgres.NewBookStorage(db),
	}
}

func (s *storage) User() repo.UserI {
	return s.userStorage
}

func (s *storage) Book() repo.BookI {
	return s.bookStorage
}
