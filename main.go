package main

import (
	"log"
	"usermanagement/router"
	"usermanagement/data"
)

func main() {
	// Initialize MongoDB
	if err := data.InitMongoDB(); err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Initialize router
	r := router.InitRouter()

	// Run server
	if err := r.Run("localhost:8000"); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
