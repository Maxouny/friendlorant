package controllers

import (
	"net/http"
	"time"

	"friendlorant/internal/models"
	"friendlorant/internal/repositories"
	"friendlorant/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userRepo repositories.UserRepository
}

func NewUserController(userRepo repositories.UserRepository) *UserController {
	return &UserController{userRepo: userRepo}
}

func (uc *UserController) Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	// field validation
	if user.Username == "" || user.Password == "" || user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
		})
		return
	}
	// hased pass
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	user.Password = hashedPassword

	if err := uc.userRepo.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	user.Token = token

	user.TokenExpire = time.Now().Add(time.Hour * 24)

	if err := uc.userRepo.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *UserController) Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	dbUser, err := uc.userRepo.GetUserByEmail(c, user.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	if err := utils.ComparePasswords(dbUser.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token, err := utils.GenerateToken(dbUser.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	dbUser.Token = token
	c.JSON(http.StatusOK, dbUser)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	// Реализация обработчика получения пользователя по ID
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	// Реализация обработчика получения пользователя по email
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	// Реализация обработчика обновления пользователя
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	// Реализация обработчика удаления пользователя
}
