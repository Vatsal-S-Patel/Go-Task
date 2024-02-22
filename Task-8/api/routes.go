package api

import (
	"fiber-mongo-api/api/userapi"
	"fiber-mongo-api/app"

	"github.com/gofiber/fiber/v2"
)

// Api struct with App to use it's DB instance
type Api struct {
	App *app.App
}

// New function will initialize Api struct and return it with error
func New(app *app.App) (*Api, error) {
	return &Api{
		App: app,
	}, nil
}

// InitializeRoutes method will initialze routes
func (api *Api) InitializeRoutes(router *fiber.App) {
	// userAPI is instance of api struct that can access all methods
	userAPI := userapi.New(api.App)

	// Grouping router with /api/user path
	userRouter := router.Group("/api/user")

	userRouter.Get("/", userAPI.GetAllUsers)
	userRouter.Post("/", userAPI.CreateUser)
	userRouter.Get("/:id", userAPI.GetUser)
	userRouter.Put("/:id", userAPI.UpdateUser)
	userRouter.Delete("/:id", userAPI.DeleteUser)
}
