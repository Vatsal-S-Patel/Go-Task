package helper

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"train-task/database"
	"train-task/model"

	"go.mongodb.org/mongo-driver/mongo"
)

// DumpDataFromCsvToDb function will read data from CSV file and dump it into Database
func DumpDataFromCsvToDb(client *mongo.Client, csvFileName string) error {
	// Open csv file
	file, err := os.Open(csvFileName)
	if err != nil {
		log.Println("ERROR: while opening csv file", err.Error())
		return err
	}
	defer file.Close()

	// make new CSV reader for that file
	csvReader := csv.NewReader(file)

	// Read first line and ignore the result of that headers
	_, _ = csvReader.Read()

	// Get Train Collection from database
	trainsCollection := database.GetCollection(client, "trains")

	// Use WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup

	for {
		// Read the next row, if csv file is not ends or any error is not occurred
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				// Wait till all goroutine executed
				wg.Wait()
				log.Println("INFO: csv file end")
				return nil
			}
			log.Println("ERROR: while reading rows from csv file", err.Error())
			return err
		}

		// Increment the WaitGroup counter
		wg.Add(1)

		// a goroutine to process the row
		go func(row []string) {
			// Defer the wg.Done()
			defer wg.Done()

			// convert trainId and trainNo into int from string
			trainId, err := strconv.Atoi(row[0])
			if err != nil {
				log.Println("ERROR: while parsing int")
				return
			}
			trainNo, err := strconv.Atoi(row[1])
			if err != nil {
				log.Println("ERROR: while parsing int")
				return
			}

			// train is the struct of Train made from CSV row
			train := model.Train{
				Sno:         trainId,
				No:          trainNo,
				Name:        row[2],
				Source:      row[3],
				Destination: row[4],
			}

			// Defining context that run on background for 10 seconds
			ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()

			// Insert that train in database
			_, err = trainsCollection.InsertOne(ctx, train)
			if err != nil {
				log.Println("ERROR: while inserting train in database", err.Error())
			}
		}(row)
	}
}
