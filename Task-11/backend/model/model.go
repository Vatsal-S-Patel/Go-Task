package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Train struct defining train with json and bson tags
type Train struct {
	Id          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	Sno         int                `json:"sno" bson:"sno"`
	No          int                `json:"trainNo" bson:"trainNo"`
	Name        string             `json:"name" bson:"name"`
	Source      string             `json:"source" bson:"source"`
	Destination string             `json:"destination" bson:"destination"`
}
