package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Contact struct {
	ID            primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	JobTitle       string             `json:"job_title" bson:"job_title"`
	Description    string             `json:"description" bson:"description"`
	Name           string             `json:"name" bson:"name"`
	CellPhone      string             `json:"cell_phone" bson:"cell_phone"`
	WorkPhone      string             `json:"work_phone" bson:"work_phone"`
	Email          string             `json:"email" bson:"email"`
	Latitude       float64            `json:"latitude" bson:"latitude"`
	Longitude      float64            `json:"longitude" bson:"longitude"`
	Street         string             `json:"street" bson:"street"`
	City           string             `json:"city" bson:"city"`
	State          string             `json:"state" bson:"state"`
	Zip            string             `json:"zip" bson:"zip"`
	CreatedDate    time.Time          `json:"created_date" bson:"created_date"`
	UpdatedDate    time.Time          `json:"updated_date" bson:"updated_date"`
	BusinessID     primitive.ObjectID `json:"business_id" bson:"business_id"`
	PersonIndex    int                `json:"person_index" bson:"person_index"`
	UserID         primitive.ObjectID `json:"user_id" bson:"user_id"`
}
