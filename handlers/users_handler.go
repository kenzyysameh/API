package handlers

import (
	"API/API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var u models.LoginReq // u feeh email w pass men el user
	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	for _, x := range models.UsersData { // x eli haylef 3la el array
		if x.Email == u.Email {
			if u.Password == x.Password {
				claims := jwt.MapClaims{
					"user_id": x.ID,
					"role":    x.Role,
					"exp":     time.Now().Add(time.Hour).Unix(),
				}

				token := jwt.NewWithClaims(
					jwt.SigningMethodHS256,
					claims,
				)

				tokenString, err := token.SignedString(
					[]byte("my-secret-key"),
				)

				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": "Could not generate token",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"token": tokenString,
				})
				return
			}
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{
		"error": "Invalid email or password",
	})
}
