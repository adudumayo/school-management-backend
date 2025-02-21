package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Main function
func main() {
	handlers.connectDB()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/learners", handlers.getLearners)
	router.POST("/learners", handlers.postLearner)
	router.GET("/learners/:id", handlers.getLearnerByID)
	router.DELETE("/learners/:id", handlers.removeLearnerByID)

	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
