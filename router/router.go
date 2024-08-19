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

	return r
}
