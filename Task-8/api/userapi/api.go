package userapi

import (
	"fiber-mongo-api/app"
	"fiber-mongo-api/app/user"
)

// api struct with app to use it's DB instance and UserService to get access of all CRUD methods defined in Service
type api struct {
	App         *app.App
	UserService user.Service
}

// New function will initialize the api struct and return it
func New(app *app.App) *api {
	return &api{
		App:         app,
		UserService: user.NewService(app),
	}
}
