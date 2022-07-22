package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type Vehicle struct {
	gorm.Model

	Plate    string `json:"plate" gorm:"typevarchar(13); uniqueindex"`
	Brand    string `json:"brand"`
	Driver   string `json:"driver"`
	Status   string `json:"status"`
	Engine   string `json:"engine"`
	KM       int    `json:"km"`
	ImageURL string `json:"imageURL"`
}

var db *gorm.DB
var err error

func main() {
	// Loading enviroment variables
	dialect := "postgres"
	host := "localhost"
	dbPort := "5432"
	user := "postgres"
	dbname := "fleetgate"
	dbpassword := "pgadmin"

	// Database connection string
	dbURI := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, dbPort, user, dbpassword, dbname)

	// Openning connection to database
	db, err = gorm.Open(dialect, dbURI)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected to database successfully")
	}

	// Close the databse connection when the main function closes
	defer db.Close()

	// Make migrations to the database if they haven't been made already
	db.AutoMigrate(&Vehicle{})

	/*----------- API routes ------------*/
	router := mux.NewRouter()

	router.HandleFunc("/vehicles", GetVehicles).Methods("GET")
	router.HandleFunc("/vehicles/{id}", GetVehicle).Methods("GET")

	router.HandleFunc("/vehicles", CreateVehicle).Methods("POST")

	router.HandleFunc("/vehicles/{id}", DeleteVehicle).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

/*-------- API Controllers --------*/

/*----- People ------*/
func GetVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var vehicle Vehicle

	db.First(&vehicle, params["id"])

	json.NewEncoder(w).Encode(&vehicle)
}

func GetVehicles(w http.ResponseWriter, r *http.Request) {
	var vehicle []Vehicle

	db.Find(&vehicle)

	json.NewEncoder(w).Encode(&vehicle)
}

func CreateVehicle(w http.ResponseWriter, r *http.Request) {
	var vehicle Vehicle
	json.NewDecoder(r.Body).Decode(&vehicle)

	createdVehicle := db.Create(&vehicle)
	err = createdVehicle.Error
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(&createdVehicle)
}

func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var vehicle Vehicle

	db.First(&vehicle, params["id"])
	db.Delete(&vehicle)

	json.NewEncoder(w).Encode(&vehicle)
}
