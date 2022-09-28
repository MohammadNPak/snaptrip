package model

import (
	"fmt"
)

// func change_price(db *gorm.DB, tickets *[]Ticket) {
// 	var rules []Rule
// 	err := db.Model(&Rule{}).
// 		Preload("Routes").
// 		Preload("Airlines").
// 		Preload("Agencies").
// 		Preload("Suppliers").
// 		Find(&rules).Error
// 	if err != nil {
// 		panic("error")
// 	}
// }

func CreateRule(rules []Rule) {
	for i, rule := range rules {
		fmt.Printf("%d%v", i, rule)

	}
}
