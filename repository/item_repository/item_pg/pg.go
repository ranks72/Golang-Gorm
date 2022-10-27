package item_pg

import (
	"golang-gorm/models"
	"golang-gorm/repository/item_repository"

	"gorm.io/gorm"
)

type itemPG struct {
	db *gorm.DB
}

func NewItemPG(db *gorm.DB) item_repository.ItemRepository {
	return &itemPG{
		db: db,
	}
}

func (u *itemPG) Create(item models.Item) (models.Item, error) {
	err := u.db.Create(&item).Error
	return item, err
}
