package router

import (
	"github.com/gin-gonic/gin"
	"usermanagement/controllers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUsersByID)
	r.POST("/users", controllers.PostUser)
	r.DELETE("/users/:id", controllers.RemoveUser)
	r.PUT("/users/:id", controllers.UpdateUser)

	// Emoji routes
	r.GET("/emojis", controllers.GetEmojis)
	r.GET("/emojis/:id", controllers.GetEmojiByID)
	r.POST("/emojis", controllers.PostEmoji)
	r.DELETE("/emojis/:id", controllers.RemoveEmoji)
	r.PUT("/emojis/:id", controllers.UpdateEmoji)

	return r
}
