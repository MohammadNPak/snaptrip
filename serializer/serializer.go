package serializer

import (
	"fmt"
	"snapptrip/model"
	"strings"

	"gorm.io/gorm"
)

type Ticket struct {
	Origin       string `json:"origin"`
	Destination  string `json:"destination"`
	Airline      string `json:"airline"`
	Agency       string `json:"agency"`
	Supplier     string `json:"supplier"`
	BasePrice    int64  `json:"basePrice"`
	Markup       int64  `json:"markup"`
	PayablePrice int64  `json:"payablePrice"`
}

type RouteSerializer struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}

type RuleSerializer struct {
	Routes      []RouteSerializer `json:"routes"`
	Airlines    []string          `json:"airlines"`
	Agencies    []string          `json:"agencies"`
	Suppliers   []string          `json:"suppliers"`
	AmountType  string            `json:"amountType"`
	AmountValue float64           `json:"amountValue"`
}

func (self RuleSerializer) Save(db *gorm.DB) error {
	rule := model.Rule{AmountType: self.AmountType, AmountValue: self.AmountValue}
	// db.Where()
	id := rule.ID
	for agency_index:=0;agency_index<len(self.Agencies);agency_index++{
		for airlines_index:=0;airlines_index<len(self.Airlines);airlines_index++{
			for route_index:=0;route_index<len(self.Routes);route_index++{
				for suplier_index:=0;suplier_index<len(self.Suppliers);suplier_index++{
					db.Exec(fmt.Sprintf(`insert into `))
				}
			}
		}
	}
	if len(self.Airlines) != 0 {
		airlines := " (VALUES (" + strings.Join(self.Airlines, "),(") + ") a(A) "
		fmt.Printf("%s", airlines)
	}
	if len(self.Suppliers) != 0 {
		suppliers := " (VALUES (" + strings.Join(self.Suppliers, "),(") + ") a(A) "
		fmt.Printf("%s", suppliers)
	}
	if len(self.Agencies) != 0 {
		agencies := " (VALUES (" + strings.Join(self.Agencies, "),(") + ") a(A) "
		fmt.Printf("%s", agencies)
	}
	fmt.Printf("%d", id)
	return nil
	// if len(self.Routes)!=0{
	// 	routes:=" (VALUES ("+strings.Join(self.Routes.Origin,")(")+") a(A) "
	// 	fmt.Printf("%s",routes)

	// }
	// fmt.Printf("%s%s%s%s",airlines,s)
	// db.Exec(`insert into rules`)
	// SELECT N AS number, L AS letter FROM
	// (VALUES (1), (2), (3)) a(N)
	// CROSS JOIN
	// (VALUES ('A'), ('B'), ('C')) b(L);

	// const valuesLen int = len(self.Routes) * len(self.Airlines) * len(self.Agencies) * len(self.Suppliers)
	// values := [valuesLen]string
	// // Model(&model.Rule{})
	// for _, route := range self.Routes {
	// 	for _, air_line := range self.Airlines {
	// 		for _, agency := range self.Agencies {
	// 			for _, supplier := range self.Suppliers {

	// 			}
	// 		}
	// 	}
	// }
	// 		rule_id
	// 		origin
	// 		destination
	// 		airlines
	// 		agencies
	// 		suppliers
	// 		amountType
	// 		amountValue
	// `)
}
