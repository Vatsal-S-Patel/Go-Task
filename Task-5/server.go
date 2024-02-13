package main

import (
	"fmt"
	"httpserver/dbconnection"
	"httpserver/handler"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {

	envMap, err := godotenv.Read(".env")
	if err != nil {
		log.Println(err)
		return
	}

	// This will connect to the database and the instance of Database is exported
	err = dbconnection.ConnectDatabase()
	// This will close the exported Database
	defer dbconnection.CloseDatabase()
	if err != nil {
		log.Println(err)
		return
	}

	// fileServer is File Server opening file server in static folder
	fileServer := http.FileServer(http.Dir("./static"))

	// Different routes are handled via multiple Handlers
	http.Handle("/", fileServer)
	http.HandleFunc("/register", handler.FormHandler)
	http.HandleFunc("/users", handler.ShowUsersHandler)

	// Start a Web Server Listening at some Port
	err = http.ListenAndServe(fmt.Sprintf(":%v", envMap["PORT"]), nil)
	if err != nil {
		log.Println(err)
		return
	}
}
