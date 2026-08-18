package handlers

import (
	"API/API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListProducts(c *gin.Context) {
	category := c.Query("category")
	if category == "" {
		c.JSON(http.StatusOK, gin.H{
			"data": models.ProductsData,
		})
		return
	}

	filtered := []models.Product{}
	for _, p := range models.ProductsData {
		if p.Category == category {
			filtered = append(filtered, p)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": filtered,
	})
}

func CreateProduct(c *gin.Context) {
	var new models.Product
	err := c.ShouldBindJSON(&new) //reads json from request body and binds it to product struct
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	new.Id = len(models.ProductsData) + 1
	models.ProductsData = append(models.ProductsData, new)
	c.JSON(http.StatusCreated, gin.H{
		"data": new, //return el new product
	})
}
func UpdateProduct(c *gin.Context) {
	id := c.Param("id") //param haygeb el parameter eli esmo id men url path
	x, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	for i, p := range models.ProductsData {
		if p.Id == x { //p da belef 3la el loop  x da id eli 5dnah men el url
			var updated models.ProductUpdateReq
			err := c.ShouldBindJSON(&updated)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Bad Request",
				})
				return
			}
			// models.ProductsData[i]
			if updated.Price != nil {
				models.ProductsData[i].Price = *updated.Price
			}
			if updated.Stock != nil {
				models.ProductsData[i].Stock = *updated.Stock
			}
			if updated.Price == nil || updated.Stock == nil {
				return
			}

			c.JSON(http.StatusCreated, gin.H{
				"data": updated,
			})
		}
	}
}
