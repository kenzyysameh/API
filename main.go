package main

import (
	"API/API/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/products", handlers.ListProducts)

	r.Run(":8080")
}
