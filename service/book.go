package service

import (
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/abdukhashimov/golang-hex-architecture/storage"
)

type BookI interface {
	GetAllBooks() ([]models.Book, error)
	CreateBook(models.BookCreate) (models.Book, error)
	UpdateBook(string, models.BookUpdate) (models.BookUpdate, error)
	DeleteBook(string) error
}

type book struct {
	db storage.StorageI
}

func NewBookService(db storage.StorageI) BookI {
	return &book{
		db: db,
	}
}

func (b *book) GetAllBooks() ([]models.Book, error) {
	res, err := b.db.Book().GetAllBooks()

	return res, err
}

func (b *book) CreateBook(req models.BookCreate) (models.Book, error) {
	res, err := b.db.Book().CreateBook(req)

	if err != nil {
		return models.Book{}, err
	}

	return res, nil
}

func (b *book) UpdateBook(id string, req models.BookUpdate) (models.BookUpdate, error) {
	res, err := b.db.Book().UpdateBook(id, req)

	return res, err
}

func (b *book) DeleteBook(id string) error {
	return b.db.Book().DeleteBook(id)
}
