package dto

type ItemRequest struct {
	Item_Code   string `json:"item_Code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_Id    int    `json:"order_id"`
}

type ItemAllRespons struct {
	Item_Code   string `json:"item_Code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
