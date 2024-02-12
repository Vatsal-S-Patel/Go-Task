package main

import (
	"fmt"
	"httpserver/dbconnection"
	"httpserver/handler"
	"log"
	"net/http"
)

func main() {

	// This will connect to the database and the instance of Database is exported
	dbconnection.ConnectDatabase()

	// fileServer is File Server opening file server in static folder
	fileServer := http.FileServer(http.Dir("./static"))

	// Different routes are handled via multiple Handlers
	http.Handle("/", fileServer)
	http.HandleFunc("/register", handler.FormHandler)
	http.HandleFunc("/users", handler.ShowUsersHandler)

	// Start a Web Server Listening at Port 8080
	fmt.Println("Starting Web Server at port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

	// This will close the exported Database
	defer func() {
		dbconnection.Database.Close()
		fmt.Println("\nConnection Closed")
	}()
}
