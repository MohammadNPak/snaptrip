package main

import (
	"flag"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"snapptrip/management"
	"snapptrip/migrations"

	// "snapptrip/model"
	"snapptrip/serializer"
)

var db *gorm.DB

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
	rules[0].Save(db)
	rules[1].Save(db)
	// model.CreateRule(rules)
	c.JSON(http.StatusOK, rules)
}

func main() {
	newdb := flag.Bool("newdb", false, "create data base ")
	loaddata := flag.Bool("loaddb", false, "load data base from csv file in ./bootcamp/data ")
	if *newdb {
		migrations.CreateDb()
	}
	db = migrations.OpenDb()
	if *newdb {
		migrations.Migrate(db)
	}
	if *loaddata {
		management.LoadData(db)
	}

	router := gin.Default()
	router.POST("/changeprice", getPrice)
	router.POST("/createrule", createRule)

	router.Run("localhost:8080")
}
