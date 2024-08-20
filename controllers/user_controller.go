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

func GetUsers(c *gin.Context) {
	var users []models.User
	cur, err := data.Collection.Find(context.TODO(), bson.D{})
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
		var user models.User
		err := cur.Decode(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    []interface{}{},
			})
			return
		}
		users = append(users, user)
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
		"data":    users,
	})
}

func PostUser(c *gin.Context) {
	var newUser models.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	newUser.ID = primitive.NewObjectID()
	newUser.CreatedDate = time.Now()
	newUser.UpdatedDate = time.Now()

	if _, err := data.Collection.InsertOne(context.TODO(), newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created",
		"data":    newUser,
	})
}

func GetUsersByID(c *gin.Context) {
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

    var user models.User
    err = data.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
    if err == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
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
		"data":    user,
	})
}


func RemoveUser(c *gin.Context) {
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

	res, err := data.Collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
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
			"message": "User not found",
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "User removed",
		"data":    map[string]interface{}{},

	})
}

func UpdateUser(c *gin.Context) {
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

    var updatedUser models.User
    if err := c.ShouldBindJSON(&updatedUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": err.Error(),
            "data":    map[string]interface{}{},
        })
        return
    }

    // Fetch the existing user to retain fields that are not being updated
    var existingUser models.User
    err = data.Collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&existingUser)
    if err == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "User not found",
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

    // Update only the fields provided, keeping other fields unchanged
    if updatedUser.Name != "" {
        existingUser.Name = updatedUser.Name
    }
    if updatedUser.Color_Code != "" {
        existingUser.Color_Code = updatedUser.Color_Code
    }
    existingUser.UpdatedDate = time.Now()

    update := bson.M{
        "$set": existingUser,
    }

    res, err := data.Collection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "status":  http.StatusInternalServerError,
            "message": err.Error(),
            "data":    map[string]interface{}{},
        })
        return
    }

    if res.MatchedCount == 0 {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "User not found",
            "data":    map[string]interface{}{},
        })
        return
    }

    // Return the updated user data
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "User updated",
        "data":    existingUser, 
    })
}
