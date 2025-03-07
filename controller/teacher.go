package controller

import (
	"net/http"

	"github.com/adudumayo/school-management-backend/model"
	"github.com/adudumayo/school-management-backend/view"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Get all quizzes
func GetQuizzes(c *gin.Context) {
	rows, err := model.DB.Query("SELECT id, subject, topic, question, due_date FROM quiz")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch quizzes"})
		return
	}
	defer rows.Close()

	var quizzes []view.Quiz
	for rows.Next() {
		var q view.Quiz
		if err := rows.Scan(&q.ID, &q.Subject, &q.Topic, &q.Question, &q.Due_date); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing quiz data"})
			return
		}
		quizzes = append(quizzes, q)
	}

	c.JSON(http.StatusOK, quizzes)
}

// Add a new quiz
func PostQuiz(c *gin.Context) {
	var newQuiz view.Quiz
	if err := c.BindJSON(&newQuiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	result, err := model.DB.Exec("INSERT INTO quiz (subject, topic, question, due_date) VALUES (?, ?, ?, ?)",
		newQuiz.Subject, newQuiz.Topic, newQuiz.Question, newQuiz.Due_date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert quiz"})
		return
	}

	id, _ := result.LastInsertId()
	newQuiz.ID = int(id)
	c.JSON(http.StatusCreated, newQuiz)
}
