package service

import (
	"golang-gorm/dto"
	"golang-gorm/models"
	"golang-gorm/repository/item_repository"
)

type ItemService interface {
	CreateItem(item dto.ItemRequest) (models.Item, error)
}

type itemService struct {
	itemRepo item_repository.ItemRepository
}

func NewItemService(itemRepo item_repository.ItemRepository) ItemService {
	return &itemService{
		itemRepo: itemRepo,
	}
}

func (i *itemService) CreateItem(item dto.ItemRequest) (models.Item, error) {
	items := models.Item{
		Item_Code:   item.Item_Code,
		Description: item.Description,
		Quantity:    item.Quantity,
		Order_Id:    item.Order_Id,
	}
	datas, err := i.itemRepo.Create(items)
	return datas, err
}
