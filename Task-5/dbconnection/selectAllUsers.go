package dbconnection

import "httpserver/model"

// SelectAllUsers returns all user from database
func SelectAllUsers() []model.User {

	// usersData will contain all users data
	var usersData []model.User

	// Storing the result into res, produced by query execution
	res, err := Database.Query(`SELECT fname, lname, dob, email, mono from "User";`)
	if err != nil {
		panic(err)
	}
	defer res.Close()

	// Iteraing over res to access rows of returned response
	for res.Next() {
		// temporary user that will appended to usersData
		var user model.User
		// Scanning values from row and store it in user
		err = res.Scan(&user.Fname, &user.Lname, &user.Dob, &user.Email, &user.Mono)
		if err != nil {
			panic(err)
		}
		user.Dob = user.Dob[:10]
		// appending user into usersData
		usersData = append(usersData, user)
	}

	return usersData
}
