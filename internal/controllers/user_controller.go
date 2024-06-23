package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"friendlorant/internal/models"
	"friendlorant/internal/repositories"
	"friendlorant/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
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
	if err := utils.ValidateUserFields(&user); err != nil {
		validationError := err.(validator.ValidationErrors)
		errors := make(map[string]string)
		for _, vavalidationError := range validationError {
			errors[vavalidationError.Field()] = vavalidationError.Error()
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":  "Invalid user fields",
			"fields": errors,
		})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user.Password = hashedPassword

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to generate token",
		})
		return
	}

	user.Token = token

	user.TokenExpire = time.Now().Add(time.Hour * 24)

	user.CreatedAt = time.Now().UTC()
	user.UpdatedAt = user.CreatedAt

	if err := uc.userRepo.CreateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	publickUser := &models.PublickUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusCreated, publickUser)
}

func (uc *UserController) Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}
	dbUser, err := uc.userRepo.GetUserByEmail(c.Request.Context(), user.Email)

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

	publickUser := &models.PublickUser{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Image:     dbUser.Image,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	c.JSON(http.StatusOK, publickUser)
}

func (uc *UserController) GetUserByID(c *gin.Context) {
	userIDParam := c.Param("id")

	userId, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	dbUser, err := uc.userRepo.GetUserByID(c.Request.Context(), uint(userId))
	if err != nil {
		fmt.Println(err)
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}
	publickUser := &models.PublickUser{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Image:     dbUser.Image,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}

	c.JSON(http.StatusOK, publickUser)
}

func (uc *UserController) GetUserByEmail(c *gin.Context) {
	userEmail := c.Param("email")
	if userEmail == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Email is required",
		})
		return
	}
	dbUser, err := uc.userRepo.GetUserByEmail(c.Request.Context(), userEmail)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}
	publickUser := &models.PublickUser{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Image:     dbUser.Image,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	c.JSON(http.StatusOK, publickUser)
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	var user models.User

	// get user id from url
	userIDParam := c.Param("id")
	userId, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	// find user by id

	existUser, err := uc.userRepo.GetUserByID(c.Request.Context(), uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}
	// bind request body to user struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload",
		})
		return
	}

	// validate user struct
	user.ID = existUser.ID

	// hash password if password is not empty
	if user.Password != "" {
		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to hash password",
			})
			return
		}
		user.Password = hashedPassword
	} else {
		user.Password = existUser.Password
	}
	// update user
	if err := uc.userRepo.UpdateUser(c.Request.Context(), &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to update user",
		})
		return
	}

	user.UpdatedAt = time.Now()

	publickUser := &models.PublickUser{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Image:     user.Image,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	c.JSON(http.StatusOK, publickUser)
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	// get user id from url
	userIDParam := c.Param("id")
	userId, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid user ID",
		})
		return
	}

	// find user by id
	err = uc.userRepo.DeleteUser(c.Request.Context(), uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully", "userId": userId})
}

func (uc *UserController) GetUserByUsername(c *gin.Context) {
	userName := c.Param("username")
	if userName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username is required",
		})
		return
	}
	dbUser, err := uc.userRepo.GetUserByUsername(c.Request.Context(), userName)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}
	publickUser := &models.PublickUser{
		ID:        dbUser.ID,
		Username:  dbUser.Username,
		Email:     dbUser.Email,
		Image:     dbUser.Image,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
	c.JSON(http.StatusOK, publickUser)
}

func (uc *UserController) GetUsers(c *gin.Context) {
	publickUser, err := uc.userRepo.GetUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("failed to get users: %v", err),
		})
		return
	}
	c.JSON(http.StatusOK, publickUser)
}
