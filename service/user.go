package service

import (
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/abdukhashimov/golang-hex-architecture/storage"
)

type UserI interface {
	CreateUser(models.User) (models.User, error)
}

type user struct {
	db storage.StorageI
}

func NewUserService(db storage.StorageI) UserI {
	return &user{
		db:db,
	}
}

func (u *user) CreateUser(req models.User) (models.User, error) {
	res, err := u.db.User().CreateUser(req)

	if err != nil {
		return models.User{}, err
	}

	return res, nil
}


