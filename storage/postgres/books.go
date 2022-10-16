package postgres

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/abdukhashimov/golang-hex-architecture/config"
	"github.com/abdukhashimov/golang-hex-architecture/service/models"
	"github.com/abdukhashimov/golang-hex-architecture/storage/repo"
	"gorm.io/gorm"
)

type books struct {
	db *gorm.DB
}

func NewBookStorage(db *gorm.DB) repo.BookI {
	return &books{db: db}
}

func (b *books) GetAllBooks() ([]models.Book, error) {
	var books []models.Book

	config.ConnectDB().Find(&books)

	return books, nil
}

func (b *books) CreateBook(payload models.BookCreate) (models.Book, error) {
	// request to get book data
	res, err := http.Get("https://openlibrary.org/isbn/" + payload.Isbn + ".json")
	if err != nil {
		return models.Book{}, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.Book{}, err
	}

	var book models.BookResponse
	bookJson := string(body)

	if err = json.Unmarshal([]byte(bookJson), &book); err != nil {
		return models.Book{}, err
	}

	// request to get author data
	var author models.BookAuthor
	if len(book.Authors) != 0 {
		res, err = http.Get("https://openlibrary.org/" + book.Authors[0].Key + ".json")
		if err != nil {
			return models.Book{}, err
		}

		defer res.Body.Close()

		body, err = ioutil.ReadAll(res.Body)
		if err != nil {
			return models.Book{}, err
		}

		authorJson := string(body)
		if err = json.Unmarshal([]byte(authorJson), &author); err != nil {
			return models.Book{}, err
		}
	}

	// cenverting string date to int
	published, _ := strconv.ParseInt(book.PublishDate[len(book.PublishDate)-4:], 10, 32)

	newBook := models.Book{
		Book: &models.BookData{
			Isbn:      payload.Isbn,
			Title:     book.Title,
			Author:    author.Name,
			Pages:     book.Pages,
			Published: uint(published),
		},
		Status: 0,
	}

	result := config.ConnectDB().Create(&newBook)

	return newBook, result.Error
}

func (b *books) UpdateBook(id string, payload models.BookUpdate) (models.BookUpdate, error) {

	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return models.BookUpdate{}, err
	}

	result := config.ConnectDB().Model(&models.Book{}).Where("id=?", uint(bookId)).Updates(&payload)
	payload.Book.ID = uint(bookId)

	return payload, result.Error
}

func (b *books) DeleteBook(id string) error {

	bookId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	result := config.ConnectDB().Where("id=?", uint(bookId)).Delete(&models.Book{})
	if result.RowsAffected < 1 {
		return errors.New("Error to deleting a book")
	}

	return nil
}
