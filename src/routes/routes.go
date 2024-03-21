package routes

import (
	"assignment2/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/orders", controllers.StoreOrder)
	route.GET("/orders", controllers.IndexOrder)
	route.GET("/orders/:id", controllers.ShowOrder)
	route.PATCH("/orders/:id", controllers.UpdateOrder)
	route.DELETE("/orders/:id", controllers.DestroyOrder)

	route.POST("/items", controllers.StoreItem)
	route.GET("/items", controllers.IndexItem)
	route.GET("/items/:id", controllers.ShowItem)
	route.PATCH("/items/:id", controllers.UpdateItem)
	route.DELETE("/items/:id", controllers.DestroyItem)

	route.Run()
}
