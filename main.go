package main

import (
	"API/API/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("api/v1/products", handlers.ListProducts)
	r.POST("api/v1/products", handlers.CreateProduct)
	r.PATCH("api/v1/products/:id", handlers.UpdateProduct)

	r.Run(":8080")

}
