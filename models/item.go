package models

type Item struct {
	Id          uint   `json:"id_item" gorm:"primary_key"`
	Item_Code   string `json:"item_code" gorm:"not null;type:varchar(255)"`
	Description string `json:"description" gorm:"not null;type:text"`
	Quantity    int    `json:"quantity" gorm:"not null;type:int4"`
	Order_Id    int    `json:"order_id" gorm:"not null"`
}
