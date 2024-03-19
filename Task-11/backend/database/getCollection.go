package database

import (
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection function will return *mongo.Collection if exist
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {

	// Returning Collection from database
	return client.Database(os.Getenv("DBNAME")).Collection(collectionName)
}
