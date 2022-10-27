package item_repository

import "golang-gorm/models"

type ItemRepository interface {
	Create(item models.Item) (models.Item, error)
}
