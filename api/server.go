package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	// Define your routes here
	// Example:
	// router.GET("/users", getUsers)
	// router.POST("/users", createUser)
	// router.GET("/users/:id", getUserByID)
	// router.PUT("/users/:id", updateUser)
	// router.DELETE("/users/:id", deleteUser)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
