package order_repository

import "golang-gorm/models"

type OrderRepository interface {
	CreateOrder(data models.Order) (models.Order, error)
	GetOrders() ([]models.Order, error)
	FindById(OrderID int) (models.Order, error)
	UpdateOrder(OrderID int, data models.Order) (models.Order, error)
	// UpdateItems(orderID int, Items []models.Item) error
	// DeleteOrder(id int) error
}
