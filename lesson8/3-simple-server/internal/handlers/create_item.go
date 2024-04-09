package handlers

import (
	"encoding/json"
	"net/http"

	"simple-server/internal/domain"
	"simple-server/internal/entities"
	"simple-server/internal/handlers/dtos"
)

type createItemHandler struct {
	itemRepo domain.ItemRepository
}

func NewCreateItemHandler(itemRepo domain.ItemRepository) http.Handler {
	return &createItemHandler{itemRepo}
}

func (h *createItemHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	itemDto := &dtos.Item{}
	if err := json.NewDecoder(r.Body).Decode(itemDto); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	newItem := entities.Item{Id: itemDto.Id, Title: itemDto.Title, Description: itemDto.Description}

	if h.itemRepo.AddItem(newItem) {
		rw.WriteHeader(http.StatusCreated)
	} else {
		rw.WriteHeader(http.StatusConflict)
	}
}
