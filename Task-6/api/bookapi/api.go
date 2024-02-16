package bookapi

import (
	"book-crud-api/app"
	"book-crud-api/app/book"
)

// api struct has two reference, one is App struct and second is Service interface for book
type api struct {
	App         *app.App
	BookService book.Service // Service is interface which force to implement all the CRUD methods and use to bind method with api
}

// New function returns the api struct with app and BookService with DB instance
func New(app *app.App) *api {
	return &api{
		App:         app,                  // the app pointer
		BookService: book.NewService(app), // the DB pointer
	}
}
