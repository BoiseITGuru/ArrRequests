package controllers

import (
	"net/http"

	"github.com/BoiseITGuru/ArrRequests/internal/auth"
	"github.com/BoiseITGuru/ArrRequests/internal/models"
	"github.com/BoiseITGuru/ArrRequests/internal/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u := &models.User{}

	u.Username = input.Username
	u.Password = input.Password

	err := saveUser(u)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "registration success"})

}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := loginCheck(u.Username, u.Password)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func saveUser(user *models.User) error {
	return services.DB.Create(&user).Error
}

func verifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func loginCheck(username string, password string) (string, error) {
	var err error

	u := models.User{}

	err = services.DB.Model(models.User{}).Where("username = ?", username).Take(&u).Error
	if err != nil {
		return "", err
	}

	err = verifyPassword(password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := auth.GenerateJWT(u)
	if err != nil {
		return "", err
	}

	return token, nil

}
