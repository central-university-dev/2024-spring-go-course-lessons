package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-framework/internal/domain"
	"gin-framework/internal/entities"
	"gin-framework/internal/handlers/dtos"
)

func SetupItemHandlers(g gin.IRoutes, itemRepo domain.ItemRepository) {
	g.GET("/item", SetupGetItemsHandler(itemRepo))
	g.GET("/item/:itemId", SetupGetItemByIdHandler(itemRepo))
	g.POST("/item", SetupCreateItemHandler(itemRepo))
}

// GetItems godoc
// @Summary      Получение списка вещей
// @Produce      json
// @Success      200  {array}  dtos.Item
// @Router       /item [get]
func SetupGetItemsHandler(itemRepo domain.ItemRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		items := itemRepo.GetItems()
		itemsDto := make([]dtos.Item, len(items))

		for i, item := range items {
			itemsDto[i] = dtos.Item{Id: item.Id, Title: item.Title, Description: item.Description}
		}

		ctx.JSON(http.StatusOK, itemsDto)
	}
}

// GetItemById godoc
// @Summary      Получение вещи по идентификатору
// @Produce      json
// @Param        itemId   path      int  true  "Идентификатор вещи"
// @Success      200  {object}  dtos.Item
// @Router       /item/{itemId} [get]
func SetupGetItemByIdHandler(itemRepo domain.ItemRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemId, err := strconv.ParseInt(ctx.Param("itemId"), 10, 64)
		if err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		item := itemRepo.GetItemById(itemId)

		if item != nil {
			itemDto := dtos.Item{Id: item.Id, Title: item.Title, Description: item.Description}
			ctx.JSON(http.StatusOK, itemDto)
		} else {
			ctx.AbortWithStatus(http.StatusNotFound)
		}
	}
}

// CreateItem godoc
// @Summary      Добавить вещь
// @Accept       json
// @Param        item body dtos.Item true "Новая вещь"
// @Success      200
// @Router       /item [post]
func SetupCreateItemHandler(itemRepo domain.ItemRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		itemDto := &dtos.Item{}
		if err := ctx.BindJSON(itemDto); err != nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		newItem := entities.Item{Id: itemDto.Id, Title: itemDto.Title, Description: itemDto.Description}

		if itemRepo.AddItem(newItem) {
			ctx.Status(http.StatusCreated)
		} else {
			ctx.AbortWithStatus(http.StatusConflict)
		}
	}
}
