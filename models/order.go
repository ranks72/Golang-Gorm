package models

import "time"

type Order struct {
	Order_Id      uint      `json:"order_id" gorm:"primary_key"`
	Customer_Name string    `json:"customer_name" gorm:"not null;type:varchar(255)"`
	Ordered_At    time.Time `json:"ordered_at" gorm:"not null; default:now()"`
	Items         Item      `json:"items" gorm:"foreignKey:Order_Id;references:Order_Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
