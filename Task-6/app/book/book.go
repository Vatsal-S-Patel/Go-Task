package book

import (
	"book-crud-api/model"
	"errors"
	"strconv"

	"gorm.io/gorm"
)

// createBook will insert the book into database
func createBook(db *gorm.DB, book model.Book) error {
	result := db.Create(&book)

	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// getAllBook return all books from database
func getAllBook(db *gorm.DB) ([]model.Book, error) {
	var books []model.Book

	err := db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return books, nil
}

// getOneBook return the first book matched with id
func getOneBook(db *gorm.DB, id string) (model.Book, error) {
	var book model.Book

	result := db.First(&book, id)
	bookId, err := strconv.Atoi(id)
	if result.Error != nil || err != nil || bookId != book.Id {
		return model.Book{}, errors.New("ERROR: Book not found for given ID")
	}

	return book, nil
}

// updateBook will update the book matched with id, and if id not found it will create new book
func updateBook(db *gorm.DB, book model.Book, id string) error {
	dbbook, err := getOneBook(db, id)
	if err != nil {
		return err
	}

	dbbook.Title = book.Title
	dbbook.Author = book.Author
	dbbook.ISBN = book.ISBN
	dbbook.Publisher = book.Publisher
	dbbook.Year = book.Year
	dbbook.Genre = book.Genre

	result := db.Save(&dbbook)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}

// deleteBook will delete the book from database matched with id
func deleteBook(db *gorm.DB, id string) error {
	var book model.Book

	// Permanently delete the book instead of soft delete
	result := db.Unscoped().Delete(&book, id)
	if result.Error != nil || result.RowsAffected == 0 {
		return result.Error
	}

	return nil
}
