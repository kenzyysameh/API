package handlers

import (
	"API/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	category :=c.Query("category")      //?category= electronics
	if category ==""  {
	c.JSON(http.StatusOK, gin.H{
		"data": models.ProductsData,
	})
	return 	
	}   
	filtered := []models.Product{}     //array fadi lel filtered products men no3 product "struct"
for _,product := range models.ProductsData {
	if product.Category==category {
		filtered=append(filtered , product)
c.JSON(http.StatusOK, gin.H{
		"data": models.filtered,

	})
}
	}      
	



