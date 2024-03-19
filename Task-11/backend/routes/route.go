package routes

import (
	"net/http"
	"train-task/handlers"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitializeRoutes(r *mux.Router, client *mongo.Client) {

	// CORSMiddleware to avoid CORS error for frontend
	r.Use(CORSMiddleware)

	// trainRouter is subrouter for /api/trains routes
	trainRouter := r.PathPrefix("/api/trains").Subrouter()

	trainRouter.HandleFunc("/", handlers.TrainHandler(client)).Methods("GET")
}

// CORSMiddleware to avoid CORS error in frontend
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow anyone to access origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
