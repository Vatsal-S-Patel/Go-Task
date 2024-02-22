package database

import (
	"context"
	"fiber-mongo-api/configs"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB function will connect with database
func ConnectDB() (*mongo.Client, error) {

	// Getting the MongoURI from .env file
	mongoURI, err := configs.GetEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}

	// context defining 10 second deadline for completion of task of connecting the database
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the database using MongoURI
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		return nil, err
	}

	// Check database connection using Ping method
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("INFO: Database Connection Established")

	return client, nil
}
