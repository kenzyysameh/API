/*
package handlers

import (

	"net/http"

	"github.com/gin-gonic/gin"

)

	func AuthenMiddleWare(c *gin.Context) {
		token := c.GetHeader("Authentication")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}

	func AuthorMiddleWare(c *gin.Context) {
		role := c.GetHeader("Role")
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Forbidden",
			})
			c.Abort()
			return
		}
		c.Next()
	}
*/
package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("my-secret-key")

func AuthenMiddleWare(c *gin.Context) {

	token := c.GetHeader("Authorization")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(token, "Bearer ")

	parsedToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !parsedToken.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		c.Abort()
		return
	}

	claims := parsedToken.Claims.(jwt.MapClaims)

	userId := claims["user_id"]
	role := claims["role"]

	c.Set("userId", userId)
	c.Set("role", role)

	c.Next()
}

func AuthorMiddleWare(c *gin.Context) {
	fmt.Println("yrdydfgdfgdfgdf")
	role, exists := c.Get("role")

	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Forbidden",
		})
		c.Abort()
		return
	}

	c.Next()
}
