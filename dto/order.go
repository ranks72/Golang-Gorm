package dto

type OrderRequest struct {
	Customer_Name string        `json:"customerName"`
	Items         []ItemRequest `json:"items"`
}
