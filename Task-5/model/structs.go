package model

// Defines the User struct
type User struct {
	Fname string `sql:"fname"`
	Lname string `sql:"lname"`
	Dob   string `sql:"dob"`
	Email string `sql:"email"`
	Mono  string `sql:"mono"`
}
