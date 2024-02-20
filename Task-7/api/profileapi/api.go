package profileapi

import (
	"jwt-go/app"
	"jwt-go/app/profile"
)

// api struct has two reference, one is App struct and second is Service interface for user
type api struct {
	App            *app.App
	ProfileService profile.Service // Service is interface which force to implement all the CRUD methods and use to bind method with api
}

// New function returns the api struct with app and ProfileService with DB instance
func New(app *app.App) *api {
	return &api{
		App:            app,                     // the app pointer
		ProfileService: profile.NewService(app), // the DB pointer
	}
}
