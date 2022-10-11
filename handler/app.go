package handler

import (
	"fmt"
	"golang-gorm/database"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func RunApp() {
	database.StartDb()
	db := database.GetDb()
	route := gin.Default()
	OrderHandler := newOrderHandler(db)
	orderRoute := route.Group("/orders")
	{
		orderRoute.GET("/", OrderHandler.GetAllOrders)
		orderRoute.POST("/", OrderHandler.AddOrders)
		orderRoute.PUT("/:orderId", OrderHandler.UpdateOrders)
		orderRoute.DELETE("/:orderId", OrderHandler.DeleteOrders)
	}
	fmt.Println("Server running on PORT =>", PORT)
	route.Run(PORT)
}
