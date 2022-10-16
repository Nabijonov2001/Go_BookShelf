package repo

import "github.com/abdukhashimov/golang-hex-architecture/service/models"

type BookI interface {
	GetAllBooks() ([]models.Book, error)
	CreateBook(payload models.BookCreate) (models.Book, error)
	UpdateBook(id string, payload models.BookUpdate) (models.BookUpdate, error)
	DeleteBook(id string) error
}
