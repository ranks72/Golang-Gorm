package handler

import (
	"fmt"
	"golang-gorm/database"
	"golang-gorm/repository/item_repository/item_pg"
	"golang-gorm/repository/order_repository/order_pg"
	"golang-gorm/service"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func RunApp() {
	database.StartDb()
	db := database.GetDb()
	route := gin.Default()

	OrderRepo := order_pg.NewOrderPG(db)
	OrderService := service.NewOrderService(OrderRepo)

	ItemRepo := item_pg.NewItemPG(db)
	ItemService := service.NewItemService(ItemRepo)

	OrderHandler := NewOrderHandler(OrderService, ItemService)

	orderRoute := route.Group("/orders")
	{
		orderRoute.GET("/", OrderHandler.GetAllOrders)
		orderRoute.POST("/", OrderHandler.AddOrders)
		//orderRoute.PUT("/:orderId", OrderHandler.UpdateOrders)
		//orderRoute.DELETE("/:orderId", OrderHandler.DeleteOrders)
	}
	fmt.Println("Server running on PORT =>", PORT)
	route.Run(PORT)
}
