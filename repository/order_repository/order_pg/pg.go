package order_pg

import (
	"golang-gorm/models"
	"golang-gorm/repository/order_repository"

	"gorm.io/gorm"
)

type orderPG struct {
	db *gorm.DB
}

func NewOrderPG(db *gorm.DB) order_repository.OrderRepository {
	return &orderPG{
		db: db,
	}
}

func (u *orderPG) CreateOrder(data models.Order) (models.Order, error) {
	err := u.db.Create(&data).Error
	return data, err
}

func (u *orderPG) GetOrders() ([]models.Order, error) {
	var orders []models.Order
	query := u.db.Preload("Items").Find(&orders).Error
	if query != nil {
		return nil, query
	}
	return orders, nil
}

func (u *orderPG) FindById(OrderID int) (models.Order, error) {
	var order models.Order
	err := u.db.First(&order, OrderID).Error
	return order, err
}

func (u *orderPG) UpdateOrder(OrderID int, data models.Order) (models.Order, error) {
	query := u.db.Where("id", OrderID).Updates(data)
	err := query.Error
	return data, err
}
