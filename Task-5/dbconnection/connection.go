package dbconnection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Database is exported to perform CRUD operations
var Database *sql.DB

// ConnectDatabase function connection the Database to particular database in postgreSQL
func ConnectDatabase() {

	var err error

	// Getting map from .env file and using godotenv package
	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err)
		return
	}

	// Connection String
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", envMap["USER"], envMap["PASSWORD"], envMap["HOST"], envMap["DBNAME"])

	// Database connection established with postgreSQL using connection string
	Database, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println("Connection Established!")
}
