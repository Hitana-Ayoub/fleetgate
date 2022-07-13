package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type vehicle struct {
	Id       int    `json:"id"`
	Plate    string `json:"plate"`
	Brand    string `json:"brand"`
	Driver   string `json:"driver"`
	Status   string `json:"status"`
	Engine   string `json:"engine"`
	Km       int    `json:"km"`
	ImageUrl string `json:"imageUrl"`
}

var vehicles = []vehicle{
	{
		Id:       1,
		Plate:    "150-TN-4578",
		Brand:    "Dacia",
		Driver:   "Imed",
		Status:   "Working",
		Engine:   "Diesel",
		Km:       452147,
		ImageUrl: "https://www.auto-plus.tn/assets/modules/newcars/dacia/logan-mcv/couverture/dacia-logan-mcv.jpg",
	},
	{
		Id:       2,
		Plate:    "170-TN-7778",
		Brand:    "Peugeout",
		Driver:   "Aymen",
		Status:   "Broken",
		Engine:   "Diesel",
		Km:       525145,
		ImageUrl: "https://www.auto-plus.tn/assets/modules/newcars/peugeot/bipper/couverture/peugeot-bipper.jpg",
	},
	{
		Id:       3,
		Plate:    "155-TN-4558",
		Brand:    "Isuzu",
		Driver:   "Houssem",
		Status:   "Working",
		Engine:   "Diesel",
		Km:       478965,
		ImageUrl: "https://www.auto-plus.tn/assets/modules/newcars/isuzu/dmax-4p/couverture/isuzu_dmax-4p.jpg",
	},
}

func getVehicles(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, vehicles)
}

func addVehicle(context *gin.Context) {
	var newVehicle vehicle

	if err := context.BindJSON(&newVehicle); err != nil {
		return
	}
	vehicles = append(vehicles, newVehicle)

	context.IndentedJSON(http.StatusCreated, newVehicle)
}

func getVehicleByID(id int) (*vehicle, error) {
	for i, v := range vehicles {
		if v.Id == id {
			return &vehicles[i], nil
		}
	}

	return nil, errors.New("vehicle not found")
}

func getVehicle(context *gin.Context) {
	id := context.Param("id")
	idint, err := strconv.Atoi(id)
	vehicle, err := getVehicleByID(idint)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "vehicle not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, vehicle)
}

func main() {

	router := gin.Default()
	router.GET("/vehicles", getVehicles)
	router.GET("/vehicles/:id", getVehicle)
	router.POST("/vehicles", addVehicle)
	router.Run("localhost:9090")

}
