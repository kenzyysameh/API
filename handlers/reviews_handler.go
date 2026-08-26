package handlers

import (
	"API/API/models"
	"API/API/repository"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ListReviews(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid product ID",
		})
		return
	}

	reviews, err := repository.ListReviewsQuery(productID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": reviews,
	})
}
func CreateReview(c *gin.Context) {

	productID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	userID, exists := c.Get("userId")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var new models.NewReview

	err = c.ShouldBindJSON(&new)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	if new.Rating < 1 || new.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Rating must be between 1 and 5",
		})
		return
	}

	creationDate := time.Now().Format("2006-01-02 15:04:05")

	review := models.Review{
		ProductID:    productID,
		UserID:       int(userID.(float64)),
		Rating:       new.Rating,
		Comment:      new.Comment,
		CreationDate: creationDate,
	}

	review, err = repository.CreateReviewQuery(review)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": review,
	})
}
