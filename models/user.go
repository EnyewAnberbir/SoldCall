package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	ColorCode   string             `json:"colorCode" bson:"colorCode"`
	CreatedDate time.Time          `json:"createdDate" bson:"createdDate"`
	UpdatedDate time.Time          `json:"updatedDate" bson:"updatedDate"`
}
