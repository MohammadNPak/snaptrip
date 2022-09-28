package migrations

import (
	"fmt"
	c "snapptrip/config"
	"snapptrip/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var config *c.Config = c.ReadConfig()

func CreateDb() {
	db_config := &config.Database
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s", db_config.Host, db_config.Port, db_config.Username, db_config.Password, db_config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", db_config.DBName)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		panic(rs.Error)
	}
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", db_config.DBName)
		if rs := db.Exec(stmt); rs.Error != nil {
			panic(rs.Error)
		}
		sql, err := db.DB()
		if err != nil {
			panic(err)
		}
		sql.Close()
	}
}

func OpenDb() *gorm.DB {
	db_config := &config.Database
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable TimeZone=%s", db_config.Host, db_config.Port, db_config.DBName, db_config.Username, db_config.Password, db_config.TimeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&model.City{},
		&model.Airline{},
		&model.Agency{},
		&model.Supplier{},
		&model.Route{},
		&model.Rule{},
		&model.RuleData{},
	)
	
}

// insert into cities (origin_id,destination_id)
// select
// 	table1.iata as origin_id,table2.iata as destination_id
// from
// 	(select iata from cities union select null)  table1
// cross join
//  	(select iata from cities union select null )  table2
// ;

// func create_route_table(db *gorm.DB) {
// 	db.Exec(`insert into routes (origin,destination)
// 		select from
// 		(select iata from cities as o_iata)
// 		cross join
// 		(select iata from cities as d_iata)
// 		`)
// }

// func create_rule_table(db *gorm.DB) {
// 	db.Exec(`CREATE TABLE IF NOT EXISTS rules2 (
// 		id integer not null,
// 		rule_id integer not null,
// 		origin varchar(3),
// 		destination varchar(3),
// 		airlines varchar(3) ,
// 		agencies varchar(100) ,
// 		suppliers varchar(100),
// 		amountType BOOLEAN ,
// 		amountValue NUMERIC() ,
// 		PRIMARY KEY (id)
// 		CONSTRAINT fk_origin
// 			FOREIGN KEY(origin)
// 				REFERENCES cities(iata)
// 				ON DELETE CASCADE,

// 		CONSTRAINT fk_destination
// 			FOREIGN KEY(destination)
// 				REFERENCES cities(iata)
// 				ON DELETE CASCADE,

// 		CONSTRAINT fk_airlines
// 			FOREIGN KEY(airlines)
// 				REFERENCES airlines(icao)
// 				ON DELETE CASCADE,

// 		CONSTRAINT fk_agencies
// 			FOREIGN KEY(agencies)
// 				REFERENCES agencies(name)
// 				ON DELETE CASCADE,

// 		CONSTRAINT fk_suppliers
// 			FOREIGN KEY(suppliers)
// 				REFERENCES suppliers(name)
// 				ON DELETE CASCADE,

// 	  )`)
// }
