package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct defining user with json and bson tag
type User struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name" validate:"required"`
	Age      uint8              `json:"age,omitempty" bson:"age" validate:"required"`
	Location string             `json:"location,omitempty" bson:"location" validate:"required"`
	Skill    []string           `json:"skills,omitempty" bson:"skills" validate:"required"`
}
