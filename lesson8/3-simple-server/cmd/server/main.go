package main

import (
	"log"
	"net/http"

	"simple-server/internal/domain"
	"simple-server/internal/handlers"
	"simple-server/internal/repositories"
)

func main() {
	itemRepo := repositories.NewItemRepository()

	mux := http.NewServeMux()

	setupRouter(mux, itemRepo)

	log.Fatal(http.ListenAndServe(":8000", mux))
}

func setupRouter(mux *http.ServeMux, itemRepo domain.ItemRepository) {
	mux.Handle("GET /api/items", handlers.NewGetItemsHandler(itemRepo))
	mux.Handle("GET /api/items/{itemId}", handlers.NewGetItemByIdHandler(itemRepo))
	mux.Handle("POST /api/items", handlers.NewCreateItemHandler(itemRepo))
}
