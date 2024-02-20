package main

import (
	"jwt-go/api"
	"jwt-go/app"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initialize app and get server_port
	app, server_port, err := app.New()
	defer app.CloseDB()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// initialize api
	api, err := api.New(app)
	if err != nil {
		log.Println(err)
		return
	}

	// Defining all routes and start server
	err = serveAPI(api, server_port)
	if err != nil {
		log.Println(err)
		return
	}
}

// serveAPI will create a new router, initialize routes and start a http server
func serveAPI(api *api.Api, port string) error {
	router := mux.NewRouter()
	api.InitializeRoutes(router)
	log.Println("Server Started at PORT:" + port)
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
