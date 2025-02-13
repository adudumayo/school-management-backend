package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type learner struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Grade   int     `json:"class"`
	Average float64 `json:"average"`
}

var learners = []learner{
	{ID: 1, Name: "Learner One", Grade: 8, Average: 34.66},
	{ID: 2, Name: "Learner Two", Grade: 10, Average: 60.89},
	{ID: 3, Name: "Learner Three", Grade: 9, Average: 79.23},
}

func getLearners(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, learners)
}

func postLearner(c *gin.Context) {
	var newLearner learner

	if err := c.BindJSON(&newLearner); err != nil {
		return
	}

	learners = append(learners, newLearner)
	c.IndentedJSON(http.StatusCreated, newLearner)
}

func getLearnerByID(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId) // convert id to int because Param() returns a string by default
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id format"})
		return
	}

	for _, currLearner := range learners {
		if currLearner.ID == id {
			c.IndentedJSON(http.StatusOK, currLearner)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "learner not found"})
}

func removeLearnerByID(c *gin.Context) {
	strId := c.Param("id")

	id, err := strconv.Atoi(strId)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid id format"})
		return
	}

	for pos, currLearner := range learners {
		if currLearner.ID == id {
			learners = append(learners[:pos], learners[pos+1:]...)
			c.IndentedJSON(http.StatusOK, learners)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "learner does not exist"})
}

func main() {
	router := gin.Default()
	router.GET("/learners", getLearners)
	router.POST("learners", postLearner)
	router.GET("/learners/:id", getLearnerByID)

	router.Run("localhost:8080")
}
