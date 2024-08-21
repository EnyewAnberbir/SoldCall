package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Business struct {
	ID                 primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	BusinessName        string             `json:"business_name" bson:"business_name"`
	EmojiID             primitive.ObjectID `json:"emoji_id" bson:"emoji_id"`
	BusinessTagline     string             `json:"business_tagline" bson:"business_tagline"`
	Website             string             `json:"website" bson:"website"`
	Status              int                `json:"status" bson:"status"`
	AutoFollowup        bool               `json:"auto_followup" bson:"auto_followup"`
	LastViewedDate      time.Time          `json:"last_viewed_date" bson:"last_viewed_date"`
	LastFollowupDate    time.Time          `json:"last_followup_date" bson:"last_followup_date"`
	NextFollowupDate    time.Time          `json:"next_followup_date" bson:"next_followup_date"`
	CreatedDate         time.Time          `json:"created_date" bson:"created_date"`
	UserID              primitive.ObjectID `json:"user_id" bson:"user_id"`
	ContactID           primitive.ObjectID `json:"contact_id" bson:"contact_id"`
}
