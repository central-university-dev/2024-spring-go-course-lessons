package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"gin-framework/internal/handlers"
	"gin-framework/internal/middleware"
	"gin-framework/internal/repositories"
)

// @title           API складского учёта
// @version         0.1.2
// @description     API с использованием фреймворка gin

// @contact.name   Иван Иваныч Иванов
// @contact.email  i.i.ivanov@example.com

// @host      localhost:8000
// @BasePath  /api
func main() {
	itemRepo := repositories.NewItemRepository()

	g := gin.Default()

	g.Use(middleware.GlobalLogger())
	apiGroup := g.Group("/api", middleware.NotFoundCounter())

	handlers.SetupItemHandlers(apiGroup, itemRepo)

	log.Fatal(g.Run(":8000"))
}
