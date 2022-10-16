package service

import (
	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/storage"
	"go.uber.org/zap"
)

type serviceHandler struct {
	userService    UserI
	bookService    BookI
	cleanupService CleanupI
}

type ServiceI interface {
	User() UserI
	Book() BookI
	Cleanup() CleanupI
}

func NewServiceHandler(cfg *config.Config, log *zap.Logger, strg storage.StorageI) ServiceI {
	return &serviceHandler{
		userService:    NewUserService(strg),
		bookService:    NewBookService(strg),
		cleanupService: NewCleanupService(strg),
	}
}

func (s *serviceHandler) User() UserI {
	return s.userService
}

func (s *serviceHandler) Book() BookI {
	return s.bookService
}

func (s *serviceHandler) Cleanup() CleanupI {
	return s.cleanupService
}
