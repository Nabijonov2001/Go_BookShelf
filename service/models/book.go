package models

type BookData struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Isbn      string `gorm:"unique" json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     uint   `json:"pages"`
	Published uint   `json:"published"`
}

type BookUpdateData struct {
	Isbn      string `gorm:"unique" json:"isbn"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Pages     uint   `json:"pages"`
	Published uint   `json:"published"`
}

type BookAuthor struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}

type Book struct {
	Book   *BookData `gorm:"embedded" json:"book"`
	Status uint      `sql:"type:ENUM(0, 1, 2)" gorm:"column:status" json:"status"`
}

type BookCreate struct {
	Isbn string `json:"isbn" binding:"required"`
}

type BookUpdate struct {
	Book   *BookData `gorm:"embedded" json:"book"`
	Status uint      `json:"status"`
}

type BookResponse struct {
	Authors     []BookAuthor `json:"authors,omitempty"`
	Title       string       `json:"title,omitempty"`
	Pages       uint         `json:"number_of_pages,omitempty"`
	PublishDate string       `json:"publish_date,omitempty"`
}
