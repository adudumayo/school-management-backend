package main

import (
	"fmt"
	"log"

	"github.com/adudumayo/school-management-backend/controller"
	"github.com/adudumayo/school-management-backend/model"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	model.ConnectDB()

	router := gin.Default()
	router.Use(cors.Default())

	// admin endpoints
	router.GET("/learners", controller.GetLearners)
	router.GET("/learners/:id", controller.GetLearnerByID)
	router.POST("/learners", controller.PostLearner)
	router.DELETE("/learners/:id", controller.DeleteLearnerByID)

	// admin teacher related endpoints
	router.POST("/teachers", controller.PostTeacher)

	// teacher endpoints
	router.GET("/quizzes", controller.GetQuizzes)
	router.POST("/quizzes", controller.PostQuiz)

	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
