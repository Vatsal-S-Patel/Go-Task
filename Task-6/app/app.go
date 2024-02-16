package app

import (
	"book-crud-api/database"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// App struct contains database pointer
type App struct {
	DB *gorm.DB
}

// New function returns the app instance contains with db pointer, server port and error
func New() (*App, string, error) {

	// Reading .env file
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	// Initializing Database
	db, err := database.InitDB(envMap)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}

	return &App{
			DB: db,
		},
		envMap["SERVER_PORT"],
		nil
}

// CloseDB function close the database and return error if any
func (a *App) CloseDB() error {

	// sqldb is Database instance from App's DB
	sqldb, err := a.DB.DB()
	if err != nil {
		return err
	}

	// Closing the Database connection
	err = sqldb.Close()
	if err != nil {
		return err
	}

	log.Println("Database Connection Closed")

	return nil
}
