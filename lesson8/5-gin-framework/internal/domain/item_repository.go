package domain

import (
	"gin-framework/internal/entities"
)

type ItemRepository interface {
	GetItems() []entities.Item
	GetItemById(int64) *entities.Item
	AddItem(entities.Item) bool
}
