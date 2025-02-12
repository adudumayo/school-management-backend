package main

import (
	"fmt"
	"net/http"
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

func main() {
	fmt.Println("vim-go")
}
