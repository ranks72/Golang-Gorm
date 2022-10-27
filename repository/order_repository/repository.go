package order_repository

import "golang-gorm/models"

type OrderRepository interface {
	CreateOrder(data models.Order) (models.Order, error)
	GetOrders() ([]models.Order, error)
	// UpdateOrder(id int, data models.Order) error
	// UpdateItems(orderID int, Items []models.Item) error
	// DeleteOrder(id int) error
}
