package userapi

import (
	"fiber-mongo-api/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Validator to validate struct
var validate = validator.New()

// CreateUser method will insert user in database
// POST /api/user
func (api *api) CreateUser(c *fiber.Ctx) error {
	var user model.User

	// Parsing Body of JSON input and store it in user
	err := c.BodyParser(&user)
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Validating user by matching it with User struct
	err = validate.Struct(&user)
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Assign new ObjectID to user
	user.Id = primitive.NewObjectID()

	// Create User in database using UserService's CreateUser method
	err = api.UserService.CreateUser(&user)
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Returning Successful Creation Reponse
	return model.CreatedResponse(c, user)
}

// GetAllUsers method will return all users from database
// GET /api/user
func (api *api) GetAllUsers(c *fiber.Ctx) error {
	// Retrieve all users from database using UserService's GetAllUser method
	users, err := api.UserService.GetAllUser()
	if err != nil {
		return model.InternalServerErrorResponse(c, err)
	}

	// Returning Ok Response
	return model.OkResponse(c, users)
}

// GetUser method will return user from database with mathcing id
// GET /api/user/{id}
func (api *api) GetUser(c *fiber.Ctx) error {
	// Fetch id from URL parameter and convert it into ObjectID
	userId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Get user with matched id using UserService's GetUser method
	user, err := api.UserService.GetUser(userId)
	if err != nil {
		return model.InternalServerErrorResponse(c, err)
	}

	// Returning Ok Response
	return model.OkResponse(c, user)
}

// UpdateUser method will update user in database with matching id
// PUT /api/user/{id}
func (api *api) UpdateUser(c *fiber.Ctx) error {
	var user *model.User

	// Parsing Body of JSON input and store it in user
	err := c.BodyParser(&user)
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Validate the user by matching it with User struct
	err = validate.Struct(user)
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Fetch id from URL parameter and convert it into ObjectID
	user.Id, err = primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Update user with matched id using UserService's UpdateUser method
	user, err = api.UserService.UpdateUser(user, user.Id)
	if err != nil {
		model.InternalServerErrorResponse(c, err)
	}

	// Returning Ok Response
	return model.OkResponse(c, user)
}

// DeleteUser method will delete user in database with matching id
// DELETE /api/user/{id}
func (api *api) DeleteUser(c *fiber.Ctx) error {
	// Fetch id from URL parameter and convert into ObjectID
	userId, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return model.BadResponse(c, err)
	}

	// Delete User with matched id
	err = api.UserService.DeleteUser(userId)
	if err != nil {
		return model.InternalServerErrorResponse(c, err)
	}

	// Returning Ok Response
	return model.OkResponse(c, "User Deleted")
}
