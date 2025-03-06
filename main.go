package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/adudumayo/school-management-backend/model"
	"github.com/adudumayo/school-management-backend/view"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// Get all learners
func getLearners(c *gin.Context) {
	rows, err := db.Query("SELECT id, name, class, average FROM learner")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch learners"})
		return
	}
	defer rows.Close()

	var learners []view.Learner
	for rows.Next() {
		var l view.Learner
		if err := rows.Scan(&l.ID, &l.Name, &l.Class, &l.Average); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing learner data"})
			return
		}
		learners = append(learners, l)
	}

	c.JSON(http.StatusOK, learners)
}

// Add a new learner
func postLearner(c *gin.Context) {
	var newLearner view.Learner
	if err := c.BindJSON(&newLearner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := db.Exec("INSERT INTO learner (name, class, average) VALUES (?, ?, ?)",
		newLearner.Name, newLearner.Class, newLearner.Average)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert learner"})
		return
	}

	id, _ := result.LastInsertId()
	newLearner.ID = int(id)
	c.JSON(http.StatusCreated, newLearner)
}

// Get a learner by ID
func getLearnerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var l view.Learner
	err = db.QueryRow("SELECT id, name, class, average FROM learner WHERE id = ?", id).
		Scan(&l.ID, &l.Name, &l.Class, &l.Average)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Learner not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	c.JSON(http.StatusOK, l)
}

// Delete a learner by ID
func removeLearnerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := db.Exec("DELETE FROM learner WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete learner"})
		return
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Learner not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Learner deleted"})
}

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

	router.GET("/learners", getLearners)
	router.POST("/learners", postLearner)
	router.GET("/learners/:id", getLearnerByID)
	router.DELETE("/learners/:id", removeLearnerByID)

	fmt.Println("Server running on http://localhost:8080")
	router.Run(":8080")
}
