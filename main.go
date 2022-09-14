package main

import (
	"net/http"
	"snapptrip/migrations"
	"snapptrip/model"

	"github.com/gin-gonic/gin"
)

func getPrice(c *gin.Context) {
	var tickets []model.Ticket

	error := c.BindJSON(&tickets)
	if error != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, tickets)
}

func createRule(c *gin.Context) {
	var rules []model.Rule
	error := c.BindJSON(&rules)
	if error != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, rules)
}

func main() {
	migrations.CreateDb()
	db := migrations.OpenDb()
	migrations.Migrate(db)
	router := gin.Default()
	router.POST("/changeprice", getPrice)
	router.POST("/createrule", createRule)

	router.Run("localhost:8080")
}
