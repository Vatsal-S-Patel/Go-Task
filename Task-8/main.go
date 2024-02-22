package main

import (
	"fiber-mongo-api/api"
	"fiber-mongo-api/app"
	"fiber-mongo-api/configs"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Read env file, Connect to database
	app, err := app.New()
	defer app.CloseDB()
	if err != nil {
		log.Println(err.Error())
		return
	}

	// Initializes routes, services and apis
	api, err := api.New(app)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// router is instance of Fiber
	router := fiber.New()
	// Initializing Routes using router
	api.InitializeRoutes(router)

	// Get server port from env file
	server_port, err := configs.GetEnv("SERVER_PORT")
	if err != nil {
		log.Println(err.Error())
		return
	}

	// start to listen on server_port
	router.Listen(":" + server_port)
}
