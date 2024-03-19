package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"train-task/configs"
	"train-task/database"
	"train-task/helper"
	"train-task/routes"

	"github.com/gorilla/mux"
)

func main() {

	// Load .env file
	configs.ReadEnv()

	// Connecting database
	client, err := database.ConnectDB()
	if err != nil {
		log.Println(err.Error())
	}
	defer database.CloseDB(client)

	// readCsv flag for Dumping CSV data into database
	readCsv := flag.Bool("readcsv", false, "Read CSV file and dump data to database")

	// csvFileName flag specifies the filename of CSV
	var csvFileName string
	flag.StringVar(&csvFileName, "csv", "", "CSV file name that is used for inserting data")
	flag.Parse()

	// if readcsv flag is mentioned then it will run
	if *readCsv {
		helper.DumpDataFromCsvToDb(client, csvFileName)
		return
	}

	// Creating new mux Router
	r := mux.NewRouter()

	// FileServer opens file server in static folder
	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/", fileServer)

	// Initialize routes
	routes.InitializeRoutes(r, client)

	log.Println("INFO: started http server on port", os.Getenv("SERVER_PORT"))
	// Start HTTP server
	err = http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r)
	if err != nil {
		log.Println("ERROR: while listening on http server", err.Error())
		return
	}
}
