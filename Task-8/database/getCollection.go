package database

import (
	"fiber-mongo-api/configs"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// GetCollection function will return *mongo.Collection if exist
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	dbname, err := configs.GetEnv("DBNAME")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	// Returning collection from database
	return client.Database(dbname).Collection(collectionName)
}
