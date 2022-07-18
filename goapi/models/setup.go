package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pgadmin"
	dbname   = "fleetgate"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	var database *gorm.DB
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	database, err = gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer database.Close()

	fmt.Println("Successfully connected!")

	database.AutoMigrate(&Vehicle{})

	DB = database

}
