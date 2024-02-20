package api

import (
	"jwt-go/api/profileapi"
	"jwt-go/api/taskapi"
	"jwt-go/app"
	"jwt-go/app/auth"

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
	// {any}API is the api instance so we can use all handlers, because all handlers are binded to api struct so we need api instance to access it
	profileAPI := profileapi.New(api.App)
	taskAPI := taskapi.New(api.App)
	auth := auth.NewService(api.App)

	// Defined all routes related to book
	// We can also create SubRouter for book routes seperately, using PathPrefix() and SubRouter()
	router.HandleFunc("/login", auth.GenerateJWTToken(profileAPI.CreateProfile)).Methods("POST")

	router.HandleFunc("/profile", profileAPI.CreateProfile).Methods("POST")
	router.HandleFunc("/profile/{id}", auth.VerifyJWTToken(profileAPI.GetProfile)).Methods("GET")
	router.HandleFunc("/profile/{id}", auth.VerifyJWTToken(profileAPI.UpdateProfile)).Methods("PUT")
	router.HandleFunc("/profile/{id}", auth.VerifyJWTToken(profileAPI.DeleteProfile)).Methods("DELETE")
	router.HandleFunc("/profile/task/{id}", auth.VerifyJWTToken(profileAPI.GetAllTask)).Methods("GET")

	router.HandleFunc("/task", auth.VerifyJWTToken(taskAPI.CreateTask)).Methods("POST")
	router.HandleFunc("/task/{id}", auth.VerifyJWTToken(taskAPI.GetTask)).Methods("GET")
	router.HandleFunc("/task/{id}", auth.VerifyJWTToken(taskAPI.UpdateTask)).Methods("PUT")
	router.HandleFunc("/task/{id}", auth.VerifyJWTToken(taskAPI.DeleteTask)).Methods("DELETE")

}
