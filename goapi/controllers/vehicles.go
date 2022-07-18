package controllers

import (
	"example/goapi/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateVehicleInput struct {
	Plate    string `json:"plate"`
	Brand    string `json:"brand"`
	Driver   string `json:"driver"`
	Status   string `json:"status"`
	Engine   string `json:"engine"`
	KM       int    `json:"km"`
	ImageURL string `json:"imageURL"`
}

type UpdateVehicleInput struct {
	Plate    string `json:"plate"`
	Brand    string `json:"brand"`
	Driver   string `json:"driver"`
	Status   string `json:"status"`
	Engine   string `json:"engine"`
	KM       int    `json:"km"`
	ImageURL string `json:"imageURL"`
}

// GET /vehicles
// Get all vehicles
func GetVehicles(context *gin.Context) {
	var vehicles []models.Vehicle
	models.DB.Find(&vehicles)

	context.JSON(http.StatusOK, gin.H{"data": vehicles})
}

// POST /vehicles
// Create new vehicle
func AddVehicle(context *gin.Context) {
	// Validate input
	var input CreateVehicleInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create vehicle
	vehicle := models.Vehicle{Plate: input.Plate, Brand: input.Brand, Driver: input.Driver, Status: input.Status, Engine: input.Engine, KM: input.KM, ImageURL: input.ImageURL}
	models.DB.Create(&vehicle)
	fmt.Println(vehicle)

	context.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// GET /vehicles/:id
// Find a vehicle
func GetVehicle(context *gin.Context) {
	var vehicle models.Vehicle

	if err := models.DB.Where("id = ?", context.Param("id")).First(&vehicle).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// PATCH /vehicles/:id
// Update a vehicle
func UpdateVehicle(context *gin.Context) {
	// Get model if exist
	var vehicle models.Vehicle
	if err := models.DB.Where("id = ?", context.Param("id")).First(&vehicle).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateVehicleInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&vehicle).Updates(input)

	context.JSON(http.StatusOK, gin.H{"data": vehicle})
}

// DELETE /vehicles/:id
// Delete a vehicle
func DeleteVehicle(context *gin.Context) {
	// Get model if exist
	var vehicle models.Vehicle
	if err := models.DB.Where("id = ?", context.Param("id")).First(&vehicle).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&vehicle)

	context.JSON(http.StatusOK, gin.H{"data": true})
}
