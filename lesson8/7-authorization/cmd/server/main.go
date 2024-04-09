package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"hello": "world",
	}))

	authorized.GET("/secrets", func(c *gin.Context) {
		rawUser, exists := c.Get(gin.AuthUserKey)
		if !exists {
			c.AbortWithStatus(http.StatusUnauthorized)
		} else {
			log.Print(rawUser.(string))
			c.Status(http.StatusOK)
		}
	})

	log.Fatal(r.Run(":8000"))
}
