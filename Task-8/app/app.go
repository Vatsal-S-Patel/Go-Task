package app

import (
	"context"
	"log"
	"time"

	"fiber-mongo-api/configs"
	"fiber-mongo-api/database"

	"go.mongodb.org/mongo-driver/mongo"
)

// App struct with DB type of *mongo.Client
type App struct {
	DB *mongo.Client
}

// New function Read .env file, connect to database and initialize App with DB instance
func New() (*App, error) {
	// Reading .env file
	err := configs.ReadEnv()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &App{
		DB: db,
	}, nil
}

// CloseDB method will close the database connection
func (app *App) CloseDB() error {
	// context defining 10 second deadline for completion of task of disconnecting database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Close database connection
	err := app.DB.Disconnect(ctx)
	if err != nil {
		return err
	}

	log.Println("INFO: Database Connection Closed")

	return nil
}
