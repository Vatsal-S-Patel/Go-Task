package book

import (
	"book-crud-api/app"
	"book-crud-api/model"

	"gorm.io/gorm"
)

// service struct contains Database pointer
type service struct {
	DB *gorm.DB
}

// Service interface with all CRUD functions signature
type Service interface {
	CreateBook(book model.Book) error
	GetAllBook() ([]model.Book, error)
	GetOneBook(id string) (model.Book, error)
	UpdateBook(book model.Book, id string) error
	DeleteBook(id string) error
}

// NewService will initialize service struct with Database pointer and returned as Service
func NewService(app *app.App) Service {
	return &service{
		DB: app.DB,
	}
}

// CreateBook used to call createBook function with database of service and book
func (s *service) CreateBook(book model.Book) error {
	return createBook(s.DB, book)
}

// GetAllBook used to call getAllBook function with database of service
func (s *service) GetAllBook() ([]model.Book, error) {
	return getAllBook(s.DB)
}

// GetOneBook used to call getOneBook function with database of service and id
func (s *service) GetOneBook(id string) (model.Book, error) {
	return getOneBook(s.DB, id)
}

// UpdateBook used to call updateBook function with database of service, book and id
func (s *service) UpdateBook(book model.Book, id string) error {
	return updateBook(s.DB, book, id)
}

// DeleteBook used to call deleteBook function with database of service and id
func (s *service) DeleteBook(id string) error {
	return deleteBook(s.DB, id)
}
