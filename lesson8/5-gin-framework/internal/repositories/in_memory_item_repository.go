package repositories

import (
	"slices"

	"gin-framework/internal/domain"
	"gin-framework/internal/entities"
)

type inMemoryItemRepository struct {
	items []entities.Item
}

func NewItemRepository() domain.ItemRepository {
	return &inMemoryItemRepository{
		items: []entities.Item{},
	}
}

func (r *inMemoryItemRepository) GetItems() []entities.Item {
	return r.items
}

func (r *inMemoryItemRepository) GetItemById(id int64) (item *entities.Item) {
	for _, i := range r.items {
		if i.Id == id {
			item = &i
			break
		}
	}

	return
}

func (r *inMemoryItemRepository) AddItem(item entities.Item) bool {
	existingIdx := slices.IndexFunc(r.items, func(i entities.Item) bool { return i.Id == item.Id })

	if existingIdx == -1 {
		r.items = append(r.items, item)
		return true
	} else {
		return false
	}
}
