package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/adudumayo/school-management-backend/controller"
	"github.com/adudumayo/school-management-backend/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// Main function
func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db = model.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())

	// admin endpoints
	router.GET("/learners", controller.GetLearners)
	router.POST("/learners", controller.PostLearner)
	router.GET("/learners/:id", controller.GetLearnerByID)
	router.DELETE("/learners/:id", controller.RemoveLearnerByID)

	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
