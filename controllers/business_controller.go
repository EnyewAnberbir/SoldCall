package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"usermanagement/data"
	"usermanagement/models"
)

// Check if a user exists in the database
func userExists(userID primitive.ObjectID) (bool, error) {
	count, err := data.UserCollection.CountDocuments(context.TODO(), bson.M{"_id": userID})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Check if a contact exists in the database
func contactExists(contactID primitive.ObjectID) (bool, error) {
	count, err := data.ContactCollection.CountDocuments(context.TODO(), bson.M{"_id": contactID})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Check if an emoji exists in the database
func emojiExists(emojiID primitive.ObjectID) (bool, error) {
	count, err := data.EmojiCollection.CountDocuments(context.TODO(), bson.M{"_id": emojiID})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetBusinesses retrieves all businesses
func GetBusinesses(c *gin.Context) {
	var businesses []models.Business
	cur, err := data.BusinessCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []interface{}{},
		})
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var business models.Business
		err := cur.Decode(&business)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    []interface{}{},
			})
			return
		}
		businesses = append(businesses, business)
	}

	if err := cur.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    []interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    businesses,
	})
}

// PostBusiness creates a new business
func PostBusiness(c *gin.Context) {
	var newBusiness models.Business

	if err := c.ShouldBindJSON(&newBusiness); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	if newBusiness.UserID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Enter UserID",
			"data":    map[string]interface{}{},
		})
		return
	}

	userExists, err := userExists(newBusiness.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}
	if !userExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Incorrect UserID",
			"data":    map[string]interface{}{},
		})
		return
	}

	if !newBusiness.EmojiID.IsZero() {
		emojiExists, err := emojiExists(newBusiness.EmojiID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    map[string]interface{}{},
			})
			return
		}
		if !emojiExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Incorrect EmojiID",
				"data":    map[string]interface{}{},
			})
			return
		}
	} else {
		newBusiness.EmojiID = primitive.NewObjectID() // Set a default value if EmojiID is not provided
	}

	if !newBusiness.ContactID.IsZero() {
		contactExists, err := contactExists(newBusiness.ContactID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    map[string]interface{}{},
			})
			return
		}
		if !contactExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Incorrect ContactID",
				"data":    map[string]interface{}{},
			})
			return
		}
	} else {
		newBusiness.ContactID = primitive.NewObjectID() // Set a default value if ContactID is not provided
	}
	if newBusiness.Status < 0 || newBusiness.Status > 9 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Status must be between 0 and 9",
			"data":    map[string]interface{}{},
		})
		return
	}
	
	if newBusiness.Status == 0 {
		newBusiness.Status = 0 // Default value if not provided
	}
	

	newBusiness.ID = primitive.NewObjectID()
	newBusiness.CreatedDate = time.Now()

	if _, err := data.BusinessCollection.InsertOne(context.TODO(), newBusiness); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Business created",
		"data":    newBusiness,
	})
}

// GetBusinessByID retrieves a business by ID
func GetBusinessByID(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID format",
			"data":    map[string]interface{}{},
		})
		return
	}

	var business models.Business
	err = data.BusinessCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&business)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Business not found",
			"data":    map[string]interface{}{},
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    business,
	})
}

// RemoveBusiness deletes a business by ID
func RemoveBusiness(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID format",
			"data":    map[string]interface{}{},
		})
		return
	}

	res, err := data.BusinessCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	if res.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Business not found",
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Business deleted",
		"data":    map[string]interface{}{},
	})
}

func UpdateBusiness(c *gin.Context) {
	id := c.Param("id")
	var updatedBusiness models.Business
	if err := c.ShouldBindJSON(&updatedBusiness); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Invalid ID format",
			"data":    map[string]interface{}{},
		})
		return
	}

	if updatedBusiness.UserID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Enter UserID",
			"data":    map[string]interface{}{},
		})
		return
	}

	userExists, err := userExists(updatedBusiness.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}
	if !userExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Incorrect UserID",
			"data":    map[string]interface{}{},
		})
		return
	}

	if !updatedBusiness.EmojiID.IsZero() {
		emojiExists, err := emojiExists(updatedBusiness.EmojiID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    map[string]interface{}{},
			})
			return
		}
		if !emojiExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Incorrect EmojiID",
				"data":    map[string]interface{}{},
			})
			return
		}
	} else {
		updatedBusiness.EmojiID = primitive.NilObjectID // Set a default value if EmojiID is not provided
	}

	if !updatedBusiness.ContactID.IsZero() {
		contactExists, err := contactExists(updatedBusiness.ContactID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    map[string]interface{}{},
			})
			return
		}
		if !contactExists {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "Incorrect ContactID",
				"data":    map[string]interface{}{},
			})
			return
		}
	} else {
		updatedBusiness.ContactID = primitive.NilObjectID // Set a default value if ContactID is not provided
	}

	if updatedBusiness.Status < 0 || updatedBusiness.Status > 9 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Status must be between 0 and 9",
			"data":    map[string]interface{}{},
		})
		return
	}

	if updatedBusiness.Status == 0 {
		updatedBusiness.Status = 0 // Default value if not provided
	}

	// Exclude _id from the update
	update := bson.M{
		"$set": bson.M{
			"user_id":     updatedBusiness.UserID,
			"emoji_id":    updatedBusiness.EmojiID,
			"contact_id":  updatedBusiness.ContactID,
			"status":      updatedBusiness.Status,
			"business_name":      updatedBusiness.BusinessName,
			"business_tagline":   updatedBusiness.BusinessTagline,
			"website":            updatedBusiness.Website,
			"auto_followup":      updatedBusiness.AutoFollowup,
			"last_viewed_date":   updatedBusiness.LastViewedDate,
			"last_followup_date": updatedBusiness.LastFollowupDate,
			"next_followup_date": updatedBusiness.NextFollowupDate,
			// Add other fields as necessary
		},
	}

	_, err = data.BusinessCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Business updated",
		"data":    updatedBusiness,
	})
}
