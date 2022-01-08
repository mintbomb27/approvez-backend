package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Campaign struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	TimeCreated int64              `json:"timeCreated"`
	CreatedBy   primitive.ObjectID `json:"createdBy"`
	Name        string             `json:"name"`
	Status      string             `json:"status"`
	CoverImage  string             `json:"coverImage"`
}
