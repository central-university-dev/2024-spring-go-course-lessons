package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var notFoundOccurancies = 0

func NotFoundCounter() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()

		if c.Writer.Status() == http.StatusNotFound {
			notFoundOccurancies++
			log.Printf("Отдали %d 404-ых статусов", notFoundOccurancies)
		}
	}
}
