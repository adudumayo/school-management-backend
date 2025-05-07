package controller

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/adudumayo/school-management-backend/model"
	"github.com/adudumayo/school-management-backend/view"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Get all learners
func GetLearners(c *gin.Context) {
	rows, err := model.DB.Query("SELECT id, name, class, average FROM learner")
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
func PostLearner(c *gin.Context) {
	var newLearner view.Learner
	if err := c.BindJSON(&newLearner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := model.DB.Exec("INSERT INTO learner (name, class, average) VALUES (?, ?, ?)",
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
func GetLearnerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var l view.Learner
	err = model.DB.QueryRow("SELECT id, name, class, average FROM learner WHERE id = ?", id).
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
func DeleteLearnerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	res, err := model.DB.Exec("DELETE FROM learner WHERE id = ?", id)
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

// Add a new teacher
func PostTeacher(c *gin.Context) {
	var newTeacher view.Teacher
	if err := c.BindJSON(&newTeacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := model.DB.Exec("INSERT INTO teacher (surname, username, title, password) VALUES (?, ?, ?, ?)",
		newTeacher.Surname, newTeacher.Username, newTeacher.Title, newTeacher.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert learner"})
		return
	}

	id, _ := result.LastInsertId()
	newTeacher.ID = int(id)
	c.JSON(http.StatusCreated, newTeacher)
}
