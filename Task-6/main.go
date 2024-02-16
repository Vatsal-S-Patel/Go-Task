package main

import (
	"book-crud-api/api"
	"book-crud-api/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Initialize Database and Read env file
	app, port, err := app.New()
	defer app.CloseDB()
	if err != nil {
		log.Println(err)
		return
	}

	api, err := api.New(app)
	if err != nil {
		log.Println(err)
		return
	}

	// Defining all routes and start server
	err = serveAPI(api, port)
	if err != nil {
		log.Println(err)
		return
	}
}

// serveAPI define all routes and starts server
func serveAPI(api *api.Api, port string) error {
	router := mux.NewRouter()
	api.InitializeRoutes(router)
	log.Println("Server Started at PORT:" + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return err
	}

	return nil
}
