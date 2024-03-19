package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"train-task/database"
	"train-task/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// TrainHandler is a function that accept the *mongo.Client and return the HandlerFunc
func TrainHandler(client *mongo.Client) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// Get the Train collection
		trainCollection := database.GetCollection(client, "trains")

		// Getting query parameter from URL
		pageQuery := r.URL.Query().Get("page")
		searchQuery := r.URL.Query().Get("search")

		// perPage is how many rows to be displayed per page
		perPage := 10
		// page is the current page value
		page := 1

		// If page query parameter is present in URL then convert it into int
		if pageQuery != "" {
			var err error
			page, err = strconv.Atoi(pageQuery)
			if err != nil {
				log.Println("ERROR: while converting page string to int")
				return
			}
		}

		// trainOptions is used to implement pagination for response
		trainOptions := options.Find()
		trainOptions.SetLimit(int64(perPage)).SetSkip(int64((page - 1) * perPage))

		// searchOptions is used to make query using some keywords to display train's data
		var searchOptions interface{}
		// is search query paramater is there in URL, then use aggregation to search it using regex
		if searchQuery != "" {
			searchOptions = bson.M{
				"$or": []bson.M{
					{"name": bson.M{"$regex": searchQuery, "$options": "i"}},
					{"source": bson.M{"$regex": searchQuery, "$options": "i"}},
					{"destination": bson.M{"$regex": searchQuery, "$options": "i"}},
				},
			}
		} else {
			searchOptions = bson.M{}
		}

		// cursor is pointing to the result documents
		cursor, err := trainCollection.Find(context.Background(), searchOptions, trainOptions)
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer cursor.Close(context.Background())

		// declare trains array and store resultant documents in trains
		var trains []model.Train
		err = cursor.All(context.Background(), &trains)
		if err != nil {
			log.Println("ERROR: while putting data into trains array", err.Error())
		}

		// send trains as JSON response
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(trains)
		if err != nil {
			log.Println("ERROR: while encoding trains json data", err.Error())
			return
		}
	}
}
