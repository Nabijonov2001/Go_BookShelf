package storage

import (
	"github.com/abdukhashimov/golang-hex-architecture/storage/postgres"
	"github.com/abdukhashimov/golang-hex-architecture/storage/repo"
	"gorm.io/gorm"
)

type storage struct {
	userStorage    repo.UserI
	bookStorage    repo.BookI
	cleanupStorage repo.CleanupI
}

type StorageI interface {
	User() repo.UserI
	Book() repo.BookI
	Cleanup() repo.CleanupI
}

func NewStorage(db *gorm.DB) StorageI {
	return &storage{
		userStorage:    postgres.NewUserStorage(db),
		bookStorage:    postgres.NewBookStorage(db),
		cleanupStorage: postgres.NewCleanupStorage(db),
	}
}

func (s *storage) User() repo.UserI {
	return s.userStorage
}

func (s *storage) Book() repo.BookI {
	return s.bookStorage
}

func (s *storage) Cleanup() repo.CleanupI {
	return s.cleanupStorage
}
