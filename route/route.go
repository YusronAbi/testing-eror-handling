package routes

import (
	"inventory-management/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/products", controllers.GetProducts)
	router.POST("/products", controllers.CreateProduct)

	router.GET("/inventory", controllers.GetInventory)
	router.POST("/inventory", controllers.CreateInventory)
	router.PUT("/inventory/:id", controllers.UpdateInventory)

	router.GET("/orders", controllers.GetOrders)
	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders/:id", controllers.GetOrder)

	return router
}
