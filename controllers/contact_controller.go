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
func contactUserExist(userID primitive.ObjectID) (bool, error) {
	count, err := data.UserCollection.CountDocuments(context.TODO(), bson.M{"_id": userID})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Check if a business exists in the database
func businessExists(businessID primitive.ObjectID) (bool, error) {
	count, err := data.BusinessCollection.CountDocuments(context.TODO(), bson.M{"_id": businessID})
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetContacts retrieves all contacts
func GetContacts(c *gin.Context) {
	var contacts []models.Contact
	cur, err := data.ContactCollection.Find(context.TODO(), bson.D{})
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
		var contact models.Contact
		err := cur.Decode(&contact)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    []interface{}{},
			})
			return
		}
		contacts = append(contacts, contact)
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
		"data":    contacts,
	})
}

// PostContact creates a new contact
func PostContact(c *gin.Context) {
	var newContact models.Contact

	if err := c.ShouldBindJSON(&newContact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	if newContact.UserID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Enter UserID",
			"data":    map[string]interface{}{},
		})
		return
	}

	contactUserExist, err := contactUserExist(newContact.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}
	if !contactUserExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "UserID does not exist",
			"data":    map[string]interface{}{},
		})
		return
	}

	if newContact.BusinessID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Enter BusinessID",
			"data":    map[string]interface{}{},
		})
		return
	}

	businessExists, err := businessExists(newContact.BusinessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}
	if !businessExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "BusinessID does not exist",
			"data":    map[string]interface{}{},
		})
		return
	}

	newContact.ID = primitive.NewObjectID()
	newContact.CreatedDate = time.Now()
	newContact.UpdatedDate = time.Now()

	if _, err := data.ContactCollection.InsertOne(context.TODO(), newContact); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Contact created",
		"data":    newContact,
	})
}

// GetContactByID retrieves a contact by ID
func GetContactByID(c *gin.Context) {
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

	var contact models.Contact
	err = data.ContactCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&contact)
	if err == mongo.ErrNoDocuments {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Contact not found",
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
		"data":    contact,
	})
}

// RemoveContact deletes a contact by ID
func RemoveContact(c *gin.Context) {
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

	res, err := data.ContactCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
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
			"message": "Contact not found",
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Contact deleted",
		"data":    map[string]interface{}{},
	})
}

// UpdateContact modifies a contact by ID
func UpdateContact(c *gin.Context) {
	id := c.Param("id")
	var updatedContact models.Contact
	if err := c.ShouldBindJSON(&updatedContact); err != nil {
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

	if updatedContact.UserID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Enter UserID",
			"data":    map[string]interface{}{},
		})
		return
	}

	contactUserExist, err := contactUserExist(updatedContact.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}
	if !contactUserExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "UserID does not exist",
			"data":    map[string]interface{}{},
		})
		return
	}

	if updatedContact.BusinessID.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Enter BusinessID",
			"data":    map[string]interface{}{},
		})
		return
	}

	businessExists, err := businessExists(updatedContact.BusinessID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}
	if !businessExists {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "BusinessID does not exist",
			"data":    map[string]interface{}{},
		})
		return
	}

	updatedContact.UpdatedDate = time.Now()
	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedContact}

	result, err := data.ContactCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Contact not found",
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Contact updated",
		"data":    updatedContact,
	})
}
