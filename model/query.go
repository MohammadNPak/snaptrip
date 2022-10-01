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

// -- begin;

// -- select id from routes
// -- where origin_id=

// insert into rule_route(rule_id,route_id)
// select (insert into rules(amount_type,amount_value) values ('FIXED',1000) returning rules.id as rule_id ), routes.id from routes
// join
// (VALUES ('THR', 'MHD'), ('MHD', 'SYZ')) AS rule_routes (origin_id,destination_id)
// on routes.origin_id=rule_routes.origin_id and routes.destination_id=rule_routes.destination_id
// ;

// -- {
// -- 	"origin": "THR",
// -- 	"destination": "MHD"
// -- },
// -- {
// -- 	"origin": null,
// -- 	"destination": "SYZ"
// -- }

// with inserted_rule(id) as (insert into rules(amount_type,amount_value) values ('FIXED',20) returning rules.id),
// origin_list(iata) as (select * from (values ('THR'),('MHD'),('') ) as t1 ) ,
// destination_list(iata) as (select * from (values ('MHD'),(''),('SYZ')) as t2 ),
// routes_list(origin_iata,destination_iata) as (select * from (values (('THR'),('MHD')),(('MHD'),(null)),((null),('SYZ'))) as t12),
// airline_list(icao) as (select * from (values ('zv'),('i3')) as t3),
// agencies_list(name) as (select * from (values ('MOGHIM_TEST_AGENCY'),('SHADI_AVARAN_SAFAR')) as t4),
// suppliers_list(name) as (select * from (values ('faranegar'),('aseman')) as t5),
// routes_id(id) as (select id from routes_list right join routes on
// 				  routes_list.origin_iata=routes.origin_id and routes_list.destination_iata=routes.destination_id)
// select * from routes_id;
