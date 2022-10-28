package service

import (
	"golang-gorm/dto"
	"golang-gorm/models"
	"golang-gorm/repository/order_repository"
	"time"
)

type OrderService interface {
	CreateOrder(data dto.OrderRequest) (models.Order, error)
	GetOrder() ([]models.Order, error)
}

type orderService struct {
	orderRepo order_repository.OrderRepository
}

func NewOrderService(orderRepo order_repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (u *orderService) CreateOrder(data dto.OrderRequest) (models.Order, error) {
	order := models.Order{
		Ordered_At:    time.Now(),
		Customer_Name: data.Customer_Name,
	}
	orders, err := u.orderRepo.CreateOrder(order)
	return orders, err
}

func (u *orderService) GetOrder() ([]models.Order, error) {
	orders, err := u.orderRepo.GetOrders()
	if err != nil {
		return nil, err
	}
	return orders, nil
}
