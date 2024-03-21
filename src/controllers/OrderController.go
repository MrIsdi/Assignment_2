package controllers

import (
	"assignment2/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderRequest struct {
	CustomerName string
}

type OrderResponse struct {
	OrderRequest
}

func StoreOrder(context *gin.Context) {
	var (
		request  OrderRequest
		response OrderResponse
	)

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{}
	order.CustomerName = request.CustomerName

	result := db.Create(&order)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	response.CustomerName = request.CustomerName
	context.JSON(http.StatusCreated, response)
}

func IndexOrder(context *gin.Context) {
	var order []models.Order

	// Querying to find order datas.
	err := db.Find(&order)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    order,
	})
}

func ShowOrder(context *gin.Context) {
	Id := context.Param("id")
	order := models.Order{}

	orderById := db.Where("id = ?", Id).First(&order)
	if orderById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    orderById,
	})
}

func UpdateOrder(context *gin.Context) {
	var (
		request  OrderRequest
		response OrderResponse
	)

	Id := context.Param("id")

	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := models.Order{}

	orderById := db.Where("id = ?", Id).First(&order)
	if orderById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "order not found"})
		return
	}

	order.CustomerName = request.CustomerName

	result := db.Save(&order)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	response.CustomerName = order.CustomerName

	context.JSON(http.StatusCreated, response)
}

func DestroyOrder(context *gin.Context) {
	order := models.Order{}

	Id := context.Param("id")

	delete := db.Where("id = ?", Id).Unscoped().Delete(&order)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    Id,
	})
}
