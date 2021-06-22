package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Sucuk struct {
	Id      int    `json:"id"`
	Page    int    `json:"page"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func Login() {
	router := gin.Default()

	router.Use(corsMiddleware())
	router.POST("/post", func(c *gin.Context) {
		var model Sucuk
		c.ShouldBindJSON(&model)

		fmt.Println(model)
		c.Status(200)
	})
	router.Run(":8080")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8082")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}
