package controllers

import (
	"assignment2/src/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemRequest struct {
	ItemCode    int32
	Description string
	Quantity    int32
	OrderID     uint
}

type ItemResponse struct {
	ItemRequest
}

func StoreItem(context *gin.Context) {
	var (
		request  ItemRequest
		response ItemResponse
	)

	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Item := models.Item{}
	Item.ItemCode = request.ItemCode
	Item.Description = request.Description
	Item.Quantity = request.Quantity
	Item.OrderID = request.OrderID

	result := db.Create(&Item)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	response.ItemCode = request.ItemCode
	response.Description = request.Description
	response.Quantity = request.Quantity
	response.OrderID = request.OrderID
	context.JSON(http.StatusCreated, response)
}

func IndexItem(context *gin.Context) {
	var Item []models.Item

	err := db.Find(&Item)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    Item,
	})
}

func ShowItem(context *gin.Context) {
	Id := context.Param("id")
	Item := models.Item{}

	ItemById := db.Where("id = ?", Id).First(&Item)
	if ItemById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    ItemById,
	})
}

func UpdateItem(context *gin.Context) {
	var (
		request  ItemRequest
		response ItemResponse
	)

	Id := context.Param("id")

	if err := context.BindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Item := models.Item{}

	ItemById := db.Where("id = ?", Id).First(&Item)
	if ItemById.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Item not found"})
		return
	}

	Item.ItemCode = request.ItemCode
	Item.Description = request.Description
	Item.Quantity = request.Quantity
	Item.OrderID = request.OrderID

	result := db.Save(&Item)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	response.ItemCode = request.ItemCode
	response.Description = request.Description
	response.Quantity = request.Quantity
	response.OrderID = request.OrderID

	context.JSON(http.StatusCreated, response)
}

func DestroyItem(context *gin.Context) {
	Item := models.Item{}

	Id := context.Param("id")

	delete := db.Where("id = ?", Id).Unscoped().Delete(&Item)
	fmt.Println(delete)

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    Id,
	})
}
