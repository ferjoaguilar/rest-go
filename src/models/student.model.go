package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Student struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Dni string `json:"dni,omitempty" bson:"dni,omitempty"`
	Degree string `json:"degree,omitempty" bson:"degree,omitempty"`
}