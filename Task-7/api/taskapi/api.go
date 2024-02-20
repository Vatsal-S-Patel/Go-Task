package taskapi

import (
	"jwt-go/app"
	"jwt-go/app/task"
)

// api struct has two reference, one is App struct and second is Service interface for user
type api struct {
	App         *app.App
	TaskService task.Service // Service is interface which force to implement all the CRUD methods and use to bind method with api
}

// New function returns the api struct with app and UserService with DB instance
func New(app *app.App) *api {
	return &api{
		App:         app,                  // the app pointer
		TaskService: task.NewService(app), // the DB pointer
	}
}
