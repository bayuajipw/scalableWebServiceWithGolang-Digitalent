package database

import (
	"assignment-2/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	username = "postgres"
	password = "masbayu28"
	port     = 5432
	dbName   = "db-go-sql"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(model.Orders{}, model.Items{})
}

func GetDB() *gorm.DB {
	return db
}
