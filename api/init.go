package api

import "github.com/gin-gonic/gin"

func ApiInit() {
	router := gin.Default()
	router.Use(corsMiddleware())

	router.POST("/login", Login)
	router.GET("/kayitlar", KayitGetir)
	router.DELETE("/kayitsil", KayitSil)
	router.POST("/yenikayit", YeniKayit)
	router.PUT("/guncelle", Guncelle)

	router.Run(":5000")
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
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
