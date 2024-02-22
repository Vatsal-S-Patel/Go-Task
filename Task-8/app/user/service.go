package user

import (
	"context"
	"errors"
	"fiber-mongo-api/app"
	"fiber-mongo-api/database"
	"fiber-mongo-api/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// service struct with DB instance of type *mongo.Client and UserCollection of type *mongo.Collection
type service struct {
	DB             *mongo.Client
	UserCollection *mongo.Collection
}

// NewService will initialize service struct and return Service
// which ensures that service implements Service interface, otherwise it throws error
func NewService(app *app.App) Service {
	return &service{
		DB:             app.DB,
		UserCollection: database.GetCollection(app.DB, "users"),
	}
}

// Service interface with all user CRUD methods signature
type Service interface {
	CreateUser(user *model.User) error
	GetAllUser() (*[]model.User, error)
	GetUser(id primitive.ObjectID) (*model.User, error)
	UpdateUser(user *model.User, id primitive.ObjectID) (*model.User, error)
	DeleteUser(id primitive.ObjectID) error
}

// CreateUser creates user in database
func (s *service) CreateUser(user *model.User) error {
	// context defining 10 second deadline for completion of task of creating of user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Inserting user in database
	_, err := s.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

// GetAllUser returns all user from database
func (s *service) GetAllUser() (*[]model.User, error) {
	// context defining 10 second deadline for completion of task of getting all user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var users []model.User

	// Getting all users from database
	result, err := s.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer result.Close(ctx)

	// Iterating over result to decode every user in result
	for result.Next(ctx) {
		var user model.User
		err := result.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, nil
}

// GetUser returns user with matched id from database
func (s *service) GetUser(id primitive.ObjectID) (*model.User, error) {
	// context defining 10 second deadline for completion of task of getting a user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user *model.User

	// Get user from database matched with id and store it in user
	s.UserCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)

	return user, nil
}

// UpdateUser updates user with matched id from database
func (s *service) UpdateUser(user *model.User, id primitive.ObjectID) (*model.User, error) {
	// context defining 10 second deadline for completion of task of updating a user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Update user in database matched with id
	res, err := s.UserCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	if res.MatchedCount != 1 || err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes user with matched id from database
func (s *service) DeleteUser(id primitive.ObjectID) error {
	// context defining 10 second deadline for completion of task of deleting a user
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Delete user from database matched with id
	res, err := s.UserCollection.DeleteOne(ctx, bson.M{"_id": id})
	if res.DeletedCount == 0 || err != nil {
		return errors.New("ERROR: No User found with given ID")
	}

	return nil
}
