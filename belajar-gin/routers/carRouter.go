package routers

import (
	"belajar-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)
	router.GET("/cars/:id", controllers.GetCar)
	router.GET("/cars", controllers.GetCars)
	router.PUT("/cars/:id", controllers.UpdateCar)
	router.DELETE("/cars/:id", controllers.DeleteCar)

	return router
}
