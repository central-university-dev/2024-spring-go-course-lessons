package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Item struct {
	Id          int64   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description,omitempty"`
}

const baseUrl = "http://localhost:8000/api/items"

func main() {
	itemToAdd := Item{Id: 2, Title: "Ложка"}

	getItems()
	addItem(itemToAdd)
	getItem(itemToAdd.Id)
}

func getItems() {
	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Printf("не удалось получить список предметов: %v", err)
		return
	}

	handleResponse(resp)
}

func addItem(item Item) {
	itemJson, _ := json.Marshal(item)

	req, err := http.NewRequest(http.MethodPost, baseUrl, bytes.NewReader(itemJson))
	if err != nil {
		log.Printf("ошибка при создании http-запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("не удалось создать предмет: %v", err)
		return
	}

	handleResponse(resp)
}

func getItem(itemId int64) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // <-- Даже если успели раньше, то всё равно вызываем.

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseUrl+"/"+strconv.FormatInt(itemId, 10), nil)
	if err != nil {
		log.Print(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("не удалось получить предмет: %v", err)
		return
	}

	handleResponse(resp)
}

func handleResponse(resp *http.Response) {
	defer resp.Body.Close() // <-- Обязательно!
	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ошибка при чтении тела ответа: %v", err)
		return
	}

	log.Printf("Получили ответ с кодом %d и телом '%s'", resp.StatusCode, string(buf))
}
