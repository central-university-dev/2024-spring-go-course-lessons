package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"simple-server/internal/domain"
	"simple-server/internal/handlers/dtos"
)

type getItemByIdHandler struct {
	itemRepo domain.ItemRepository
}

func NewGetItemByIdHandler(itemRepo domain.ItemRepository) http.Handler {
	return &getItemByIdHandler{itemRepo}
}

func (h *getItemByIdHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	itemId, err := strconv.ParseInt(r.PathValue("itemId"), 10, 64)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	item := h.itemRepo.GetItemById(itemId)

	if item != nil {
		itemDto := dtos.Item{Id: item.Id, Title: item.Title, Description: item.Description}
		json.NewEncoder(rw).Encode(itemDto)
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}
