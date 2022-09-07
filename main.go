package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getPrice(c *gin.Context) {
	var tickets []Ticket

	error := c.BindJSON(&tickets)
	if error != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, tickets)
}

func createRule(c *gin.Context) {
	var rules []Rule
	error := c.BindJSON(&rules)
	if error != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, rules)
}

func main() {
	router := gin.Default()
	router.POST("/changeprice", getPrice)
	router.POST("/createrule", createRule)
	router.Run("localhost:8080")
}
