package handler

import (
	"html/template"
	"httpserver/dbconnection"
	"httpserver/model"
	"log"
	"net/http"
)

// FormHandler is used to register the user
func FormHandler(w http.ResponseWriter, r *http.Request) {

	// if path is not matched then it will return from this block
	if r.URL.Path != "/register" {
		http.Error(w, "404! Page Not Found.", http.StatusNotFound)
		return
	}

	// If request method is not POST then it will return from this block
	if r.Method != "POST" {
		http.Error(w, "Method is Not Found!", http.StatusNotFound)
		return
	}

	// Parsing the Form to extract value from it
	err := r.ParseForm()
	if err != nil {
		log.Printf("ParseForm() err %v\n", err)
		return
	}

	// user will made by extracting the values from Form using name attribute's value of input tag
	user := model.User{
		Fname: r.FormValue("first-name"),
		Lname: r.FormValue("last-name"),
		Dob:   r.FormValue("date-of-birth"),
		Email: r.FormValue("email-id"),
		Mono:  r.FormValue("mono"),
	}

	// This will insert the user into the databse
	err = dbconnection.InsertUser(user)
	if err != nil {
		log.Println(err)
		return
	}

	// Redirecting to success page if user successfully added into database
	http.Redirect(w, r, "/success.html", http.StatusSeeOther)
}

// ShowUserHandler is used to display all user
func ShowUsersHandler(w http.ResponseWriter, r *http.Request) {

	// usersData receives all user from database
	var usersData []model.User
	usersData, err := dbconnection.SelectAllUsers()
	if err != nil {
		log.Println(err)
		return
	}

	// Parsing the userdata.html file and stored in t
	t, err := template.ParseFiles("./static/userdata.html")
	if err != nil {
		log.Println(err)
		return
	}

	// Writing the modified html file to writer with passed data
	err = t.Execute(w, usersData)
	if err != nil {
		log.Println(err)
		return
	}
}
