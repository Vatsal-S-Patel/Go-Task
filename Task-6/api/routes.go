package api

import (
	"book-crud-api/api/bookapi"
	"book-crud-api/app"

	"github.com/gorilla/mux"
)

// Api struct store the app pointer
type Api struct {
	App *app.App
}

// New return Api instance with app pointer
func New(app *app.App) (*Api, error) {
	return &Api{
		App: app,
	}, nil
}

// InitializeRoutes will handle all routes using mux router
func (api *Api) InitializeRoutes(router *mux.Router) {
	// bookAPI is the api instance so we can use all handlers, because all handlers are binded to api struct so we need api instance to access it
	bookAPI := bookapi.New(api.App)

	// Defined all routes related to book
	// We can also create SubRouter for book routes seperately, using PathPrefix() and SubRouter()
	router.HandleFunc("/book", bookAPI.CreateBook).Methods("POST")
	router.HandleFunc("/books", bookAPI.GetAllBook).Methods("GET")
	router.HandleFunc("/book/{id}", bookAPI.GetOneBook).Methods("GET")
	router.HandleFunc("/book/{id}", bookAPI.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", bookAPI.DeleteBook).Methods("DELETE")
}
