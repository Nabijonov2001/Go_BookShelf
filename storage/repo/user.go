package repo

import "github.com/abdukhashimov/golang-hex-architecture/service/models"

type UserI interface {
	CreateUser(payload models.User) (models.User, error)
}
