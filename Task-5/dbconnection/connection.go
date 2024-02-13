package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type application struct {
	database *sql.DB
}

var app application

// ConnectDatabase function connection the Database to particular database in postgreSQL
func ConnectDatabase(envMap map[string]string) error {

	var err error

	// Connection String
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", envMap["USER"], envMap["PASSWORD"], envMap["HOST"], envMap["DBNAME"])

	// Database connection established with postgreSQL using connection string
	app.database, err = sql.Open("postgres", connStr)
	log.Println("Connection Established!")
	if err != nil {
		return err
	}

	return nil
}

func CloseDatabase() {
	err := app.database.Close()
	log.Println("Database Connection Closed!")
	if err != nil {
		log.Println(err)
		return
	}
}
