package main

import (
	"oauth2_api/adapters"
	"oauth2_api/presentation"

	"github.com/gin-gonic/gin"
)

func injectionMiddleware(key string, value any) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(key, value)
		ctx.Next()
	}
}

func main() {
	server := gin.Default()
	server.Use(injectionMiddleware("tokenReader", adapters.Reader{}))
	server.Use(injectionMiddleware("tokenSaver", adapters.StoreToken))
	// server.POST("/token", presentation.Introduce)
	server.GET("/token/:id", presentation.Introspect)
	server.Run(":8085")
}
