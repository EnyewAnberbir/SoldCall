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

func GetEmojis(c *gin.Context) {
	var emojis []models.Emoji
	cur, err := data.EmojiCollection.Find(context.TODO(), bson.D{})
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
		var emoji models.Emoji
		err := cur.Decode(&emoji)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    []interface{}{},
			})
			return
		}
		emojis = append(emojis, emoji)
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
		"data":    emojis,
	})
}

func PostEmoji(c *gin.Context) {
	var newEmoji models.Emoji

	if err := c.ShouldBindJSON(&newEmoji); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	newEmoji.ID = primitive.NewObjectID()
	newEmoji.Created_Date = time.Now()
    count, err := data.EmojiCollection.CountDocuments(context.TODO(), bson.D{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}
	newEmoji.Emoji_Index = int(count) + 1
	if _, err := data.EmojiCollection.InsertOne(context.TODO(), newEmoji); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Emoji created",
		"data":    newEmoji,
	})
}

func GetEmojiByID(c *gin.Context) {
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

    var emoji models.Emoji
    err = data.EmojiCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&emoji)
    if err == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Emoji not found",
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
		"data":    emoji,
	})
}

func RemoveEmoji(c *gin.Context) {
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

	res, err := data.EmojiCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
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
			"message": "Emoji not found",
			"data":    map[string]interface{}{},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Emoji removed",
		"data":    map[string]interface{}{},
	})
}

func UpdateEmoji(c *gin.Context) {
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

    var updatedEmoji models.Emoji
    if err := c.ShouldBindJSON(&updatedEmoji); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "status":  http.StatusBadRequest,
            "message": err.Error(),
            "data":    map[string]interface{}{},
        })
        return
    }

    // Fetch the existing emoji to retain fields that are not being updated
    var existingEmoji models.Emoji
    err = data.EmojiCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&existingEmoji)
    if err == mongo.ErrNoDocuments {
        c.JSON(http.StatusNotFound, gin.H{
            "status":  http.StatusNotFound,
            "message": "Emoji not found",
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
    if updatedEmoji.Emoji != "" {
        existingEmoji.Emoji = updatedEmoji.Emoji
    }
    if updatedEmoji.Emoji_Name != "" {
        existingEmoji.Emoji_Name = updatedEmoji.Emoji_Name
    }
    if updatedEmoji.Emoji_Index != 0 {
        existingEmoji.Emoji_Index = updatedEmoji.Emoji_Index
    }
    existingEmoji.Created_Date = updatedEmoji.Created_Date // or keep it unchanged if needed

    update := bson.M{
        "$set": existingEmoji,
    }

    res, err := data.EmojiCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
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
            "message": "Emoji not found",
            "data":    map[string]interface{}{},
        })
        return
    }

    // Return the updated emoji data
    c.JSON(http.StatusOK, gin.H{
        "status":  http.StatusOK,
        "message": "Emoji updated",
        "data":    existingEmoji,
    })
}
