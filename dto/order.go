package dto

import (
	"golang-gorm/models"
	"time"
)

type OrderRequest struct {
	Customer_Name string        `json:"customerName"`
	Items         []ItemRequest `json:"items"`
}

type OrderGetAllResponse struct {
	Order_Id      uint           `json:"order_id"`
	Customer_Name string         `json:"customer_name"`
	Ordered_At    time.Time      `json:"ordered_at"`
	Items         ItemAllRespons `json:"items"`
}

func ObjectAllorders(data models.Order) OrderGetAllResponse {
	return OrderGetAllResponse{
		Order_Id:      data.Order_Id,
		Customer_Name: data.Customer_Name,
		Ordered_At:    data.Ordered_At,
		Items: ItemAllRespons{
			Item_Code:   data.Items.Item_Code,
			Description: data.Items.Description,
			Quantity:    data.Items.Quantity,
		},
	}
}

func GetAllOrdersResponse(res []models.Order) (responses []OrderGetAllResponse) {
	for _, order := range res {
		responses = append(responses, ObjectAllorders(order))
	}
	return
}

type UpdateOrderRequest struct {
	Customer_Name string        `json:"customerName"`
	Items         []ItemRequest `json:"items"`
}
