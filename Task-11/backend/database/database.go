package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB will make connection with database
func ConnectDB() (*mongo.Client, error) {
	// Defining context which will run in background for 10 second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the Mongo database using MongoURI from .env file
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Println("ERROR: while connecting database", err.Error())
		return nil, err
	}

	// Check the connection by Ping to database
	if err = client.Ping(ctx, nil); err != nil {
		log.Println("ERROR: error while ping database", err.Error())
		return nil, err
	}

	log.Println("INFO: database connection established")

	return client, nil
}

// CloseDB will close the database collection
func CloseDB(client *mongo.Client) error {
	// Defining context which will run in background for 10 second
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Disconnect to the database
	err := client.Disconnect(ctx)
	if err != nil {
		log.Println("ERROR: while disconnecting database")
		return err
	}

	log.Println("INFO: database connection closed")

	return nil
}
