package main

import (
	"example/go-calculator/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CalculationResponse struct {
	Action 	string 		`json:"action"`
	X 		float64 	`json:"x"`
	Y 		float64		`json:"y"`
	Answer 	float64		`json:"answer"`
	Cached 	bool		`json:"cached"`
}

func main() {
	router := gin.Default()

	router.GET("/add/:x/:y", add)
	router.GET("/subtract/:x/:y", subtract)
	router.GET("/multiply/:x/:y", multiply)
	router.GET("/divide/:x/:y", divide)

	router.Run("localhost:8080")

}

func add(c *gin.Context) {
	// Parse input parameters from HTTP request
	x, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "X is required and it's need to be a number" });
		return
	}

	y, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "Y is required and it's need to be a number" });
		return
	}

	// Call business logic
	response, error := services.Add(x, y)
	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": error.Error() });
		return
	}

	// Return result
	c.IndentedJSON(http.StatusOK, response)
}

func subtract(c *gin.Context) {
	// Parse input parameters from HTTP request
	x, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "X is required and it's need to be a number" });
		return
	}

	y, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "Y is required and it's need to be a number" });
		return
	}

	// Call business logic
	response, error := services.Subtract(x, y)
	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": error.Error() });
		return
	}

	// Return result
	c.IndentedJSON(http.StatusOK, response)
}


func multiply(c *gin.Context) {
	// Parse input parameters from HTTP request
	x, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "X is required and it's need to be a number" });
		return
	}

	y, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "Y is required and it's need to be a number" });
		return
	}

	// Call business logic
	response, error := services.Multiply(x, y)
	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": error.Error() });
		return
	}

	// Return result
	c.IndentedJSON(http.StatusOK, response)
}

func divide(c *gin.Context) {
	// Parse input parameters from HTTP request
	x, err := strconv.ParseFloat(c.Param("x"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "X is required and it's need to be a number" });
		return
	}

	y, err := strconv.ParseFloat(c.Param("y"), 64)
	if err != nil {
		log.Println("Error: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": "Y is required and it's need to be a number" });
		return
	}

	// Call business logic
	response, error := services.Divide(x, y)
	if error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ "message": error.Error() });
		return
	}

	// Return result
	c.IndentedJSON(http.StatusOK, response)
}