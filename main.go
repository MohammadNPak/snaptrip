package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"snapptrip/management"
	"snapptrip/migrations"

	// "snapptrip/model"
	"snapptrip/serializer"
)

func getPrice(c *gin.Context) {
	var tickets []serializer.Ticket

	error := c.BindJSON(&tickets)
	if error != nil {
		return
	}
	c.IndentedJSON(http.StatusOK, tickets)
}

func createRule(c *gin.Context) {
	var rules []serializer.RuleSerializer
	error := c.BindJSON(&rules)

	if error != nil {
		return
	}
	// model.CreateRule(rules)
	c.JSON(http.StatusOK, rules)
}

func main() {
	migrations.CreateDb()
	db := migrations.OpenDb()
	migrations.Migrate(db)
	management.LoadData(db)
	router := gin.Default()
	router.POST("/changeprice", getPrice)
	router.POST("/createrule", createRule)

	router.Run("localhost:8080")
}
