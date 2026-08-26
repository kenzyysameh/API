package main

import (
	"API/API/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/api/v1/users", handlers.CreateUser)
	r.POST("api/v1/login", handlers.Login)
	r.GET("api/v1/products", handlers.ListProducts)
	r.POST("/api/v1/products", handlers.AuthenMiddleWare, handlers.AuthorMiddleWare, handlers.CreateProduct)
	r.PATCH("api/v1/products/:id", handlers.AuthenMiddleWare, handlers.AuthorMiddleWare, handlers.UpdateProduct)
	r.GET("api/v1/products/:id/reviews", handlers.ListReviews)
	r.POST("/api/v1/products/:id/reviews", handlers.AuthenMiddleWare, handlers.AuthorMiddleWare, handlers.CreateReview)
	r.POST("/api/v1/users/:userID/orders", handlers.AuthenMiddleWare, handlers.CreateOrder)
	r.DELETE("api/v1/users/:userID/orders/:orderID", handlers.AuthenMiddleWare, handlers.DeleteOrder)
	r.Run(":8080")

}
