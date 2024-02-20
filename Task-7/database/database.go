package database

import (
	"fmt"
	"log"

	"jwt-go/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB function establish the database connection and return the database pointer and error if any
func InitDB(envMap map[string]string) (*gorm.DB, error) {

	// ConnectionString to connect the database
	var connStr string = fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", envMap["USER"], envMap["PASSWORD"], envMap["HOST"], envMap["DBNAME"])

	// connection is Database instance get from gorm's Open function
	connection, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// sqldb is *sql.DB type database instance
	sqldb, err := connection.DB()
	if err != nil {
		return nil, err
	}

	// Checking the database connection using Ping() method
	err = sqldb.Ping()
	if err != nil {
		return nil, err
	}

	// Create Profile table if not exist in database
	err = connection.AutoMigrate(&model.Profile{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// Create Task table if not exist in database
	err = connection.AutoMigrate(&model.Task{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Database Connection Established")

	// return the *gorm.DB instance
	return connection, nil
}
