package handlers

import (
	"encoding/json"
	"net/http"

	"simple-server/internal/domain"
	"simple-server/internal/handlers/dtos"
)

type getItemsHandler struct {
	itemRepo domain.ItemRepository
}

func NewGetItemsHandler(itemRepo domain.ItemRepository) http.Handler {
	return &getItemsHandler{itemRepo}
}

func (h *getItemsHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	items := h.itemRepo.GetItems()
	itemsDto := make([]dtos.Item, len(items))

	for i, item := range items {
		itemsDto[i] = dtos.Item{Id: item.Id, Title: item.Title, Description: item.Description}
	}

	json.NewEncoder(rw).Encode(itemsDto)
}
