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
		&model.Agencie{},
		&model.Supplier{},
		&model.Route{},
		&model.Rule{},
		&model.Ticket{})

}
