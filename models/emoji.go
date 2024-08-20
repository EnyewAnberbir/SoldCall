package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Emoji struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Emoji        string             `json:"emoji" bson:"emoji"`
	Emoji_Name   string             `json:"emoji_name" bson:"emoji_name"`
	Emoji_Index   int          `json:"emoji_index" bson:"emoji_index"`
	Created_Date time.Time          `json:"created_date" bson:"created_date"`
}
