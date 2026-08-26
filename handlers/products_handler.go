package handlers

import (
	"API/API/models"
	"API/API/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {

	products, err := repository.GetProductsQuery()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
	})
}

func CreateProduct(c *gin.Context) {

	var newProduct models.Product

	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	if newProduct.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price cannot be negative",
		})
		return
	}

	if newProduct.Stock < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Stock cannot be negative",
		})
		return
	}

	id, err := repository.CreateProductQuery(newProduct)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Error",
		})
		return
	}

	newProduct.Id = id

	c.JSON(http.StatusCreated, gin.H{
		"data": newProduct,
	})
}
func UpdateProduct(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	var update models.ProductUpdateReq

	err = c.ShouldBindJSON(&update)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	product, err := repository.UpdateProductQuery(id, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": product,
	})
}
