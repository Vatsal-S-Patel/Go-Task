package bookapi

import (
	"book-crud-api/model"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CreateBook create the book
// POST /book
func (api *api) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.Write([]byte("ERROR: Error while decoding JSON data"))
	}

	err = api.BookService.CreateBook(book)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// GetAllBook returns all books
// GET /books
func (api *api) GetAllBook(w http.ResponseWriter, r *http.Request) {
	var books []model.Book

	books, err := api.BookService.GetAllBook()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// GetOneBook return the book according to passed id
// GET /book/{id}
func (api *api) GetOneBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	book, err := api.BookService.GetOneBook(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates the book using PUT method according to passed id
// PUT /book/{id}
func (api *api) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.Write([]byte("ERROR: Error while decoding JSON data"))
	}

	id := mux.Vars(r)["id"]

	err = api.BookService.UpdateBook(book, id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

// DeleteBook delete the book according to passed id
// DELETE /book/{id}
func (api *api) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	err := api.BookService.DeleteBook(id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Book Deleted with ID: " + id))
}
