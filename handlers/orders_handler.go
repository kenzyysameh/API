package handlers

import (
	"API/API/models"
	"API/API/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {

	userID := c.Param("userID")

	user_ID, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	var new []models.OrderItem

	err = c.ShouldBindJSON(&new)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	order := models.Order{
		UserID: user_ID,
		Items:  new,
		Status: "loading order",
	}

	order, err = repository.CreateOrderQuery(order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Header(
		"Location",
		fmt.Sprintf(
			"/api/v1/users/%d/orders/%d",
			order.UserID,
			order.ID,
		),
	)

	c.JSON(http.StatusCreated, order)
}

func DeleteOrder(c *gin.Context) {

	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	orderID, err := strconv.Atoi(c.Param("orderID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	order, err := repository.DeleteOrderQuery(userID, orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, order)
}
