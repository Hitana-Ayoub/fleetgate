package main

import (
	"example/goapi/controllers"
	"example/goapi/models"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	routerVehicles := router.Group("/vehicles")
	routerVehicles.GET("", controllers.GetVehicles)
	routerVehicles.GET("/:id", controllers.GetVehicle)
	routerVehicles.POST("", controllers.AddVehicle)
	routerVehicles.PATCH("/:id", controllers.UpdateVehicle)
	routerVehicles.DELETE("/:id", controllers.DeleteVehicle)

	models.ConnectDatabase()

	router.Run("localhost:9090")

}
