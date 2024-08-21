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

	// Contact routes
	r.GET("/contacts", controllers.GetContacts)
	r.POST("/contacts", controllers.PostContact)
	r.GET("/contacts/:id", controllers.GetContactByID)
	r.PUT("/contacts/:id", controllers.UpdateContact)
	r.DELETE("/contacts/:id", controllers.RemoveContact)

	// Business routes
	r.GET("/businesses", controllers.GetBusinesses)
	r.POST("/businesses", controllers.PostBusiness)
	r.GET("/businesses/:id", controllers.GetBusinessByID)
	r.PUT("/businesses/:id", controllers. UpdateBusiness)
	r.DELETE("/businesses/:id", controllers.RemoveBusiness)

	return r
}
