package controller

import (
	_ "github.com/go-sql-driver/mysql"
)

/*
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
*/
