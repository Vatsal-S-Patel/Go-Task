package dbconnection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// All configuration attributes
const (
	USER     = "vatsal"
	PASSWORD = "vspatel3003"
	HOST     = "localhost"
	DBNAME   = "task5"
)

// Database is exported to perform CRUD operations
var Database *sql.DB

// ConnectDatabase function connection the Database to particular database in postgreSQL
func ConnectDatabase() {

	var err error

	// Connection String
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", USER, PASSWORD, HOST, DBNAME)

	// Database connection established with postgreSQL using connection string
	Database, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection Established!")

}
