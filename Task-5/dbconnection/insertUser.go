package dbconnection

import (
	"httpserver/model"
)

// InsertUser function insert passed user into database
func InsertUser(user model.User) error {

	// Inserting user into database
	_, err := Database.Exec(`INSERT INTO "User" (fname, lname, dob, email, mono) values ($1,$2,$3,$4,$5);`, user.Fname, user.Lname, user.Dob, user.Email, user.Mono)

	return err
}
