package handlers

import (
	"API/API/models"
	"API/API/repository"
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var u models.LoginReq

	err := c.ShouldBindJSON(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}

	db, err := repository.GetData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Error",
		})
		return
	}
	defer db.Close()

	var user models.User

	err = db.QueryRow(`
		SELECT id, name, email, password, role
		FROM Users
		WHERE email = ?
	`, u.Email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(u.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
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
}

func CreateUser(c *gin.Context) {

	var newUser models.User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := repository.CreateUserQuery(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal Error",
		})
		return
	}

	newUser.ID = id

	c.JSON(http.StatusCreated, gin.H{
		"data": newUser,
	})
}
