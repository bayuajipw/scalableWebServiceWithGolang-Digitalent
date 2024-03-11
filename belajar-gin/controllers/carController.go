package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarID int    `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var CarDatas = []Car{}

func CreateCar(ctx *gin.Context) {
	newCar := Car{
		CarID: len(CarDatas) + 1,
	}

	err := ctx.ShouldBindJSON(&newCar)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	CarDatas = append(CarDatas, newCar)
	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
	})
}

func GetCar(ctx *gin.Context) {
	carID := ctx.Param("id")
	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	carData := Car{}
	for _, car := range CarDatas {
		if car.CarID == carIDInt {
			carData = car
		}
	}

	if carData.CarID == 0 {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": fmt.Sprintf("Car with id %v not found", carID),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"car": carData,
	})
}

func GetCars(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"cars": CarDatas,
	})
}

func UpdateCar(ctx *gin.Context) {
	carID := ctx.Param("id")
	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	carData := Car{
		CarID: carIDInt,
	}
	err = ctx.ShouldBindJSON(&carData)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isFound := false
	updatedCarIdx := -1
	for i, car := range CarDatas {
		if car.CarID == carIDInt {
			isFound = true
			updatedCarIdx = i
		}
	}

	if !isFound {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error":   "Data not found",
			"message": fmt.Sprintf("car with id %v not found", carIDInt),
		})
		return
	}

	CarDatas[updatedCarIdx] = carData

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully updated", carIDInt),
	})
}

func DeleteCar(ctx *gin.Context) {
	carID := ctx.Param("id")
	carIDInt, err := strconv.Atoi(carID)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	isFound := false
	deletedCarIdx := -1
	for i, car := range CarDatas {
		if car.CarID == carIDInt {
			isFound = true
			deletedCarIdx = i
		}
	}

	if !isFound {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	copy(CarDatas[deletedCarIdx:], CarDatas[deletedCarIdx+1:])
	CarDatas[len(CarDatas)-1] = Car{}
	CarDatas = CarDatas[:len(CarDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("car with id %v has been successfully deleted", carIDInt),
	})
}
