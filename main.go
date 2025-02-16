package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type learner struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Class   int     `json:"class"`
	Average float64 `json:"average"`
}

var learners = []learner{
	{ID: 1, Name: "Learner One", Class: 8, Average: 34.66},
	{ID: 2, Name: "Learner Two", Class: 10, Average: 60.89},
	{ID: 3, Name: "Learner Three", Class: 9, Average: 79.23},
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

var db *sql.DB

func main() {
	// Capture connection properties.
	/*connProps := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "myschool",
	}*/
	// Get a database handle.
	var err error
	//db, err = sql.Open("mysql", connProps.FormatDSN())
	db, err := sql.Open("mysql", "root:Admin@/myschool")
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected to the database!!")

	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/learners", getLearners)
	router.POST("learners", postLearner)
	router.GET("/learners/:id", getLearnerByID)
	router.DELETE("/learners/:id", removeLearnerByID)

	router.Run("localhost:8080")
}
