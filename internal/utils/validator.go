package utils

import (
	"context"
	"errors"

	"friendlorant/internal/models"
	"friendlorant/internal/repositories"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUserFields(user *models.User) error {
	return validate.Struct(user)
}

func ValidateCreateUser(user *models.User, userRepo repositories.UserRepository) error {
	existingUser, err := userRepo.GetUserByEmail(context.Background(), user.Email)
	if err != nil && existingUser != nil {
		return errors.New("email already exists")
	}

	existingUser, err = userRepo.GetUserByUsername(context.Background(), user.Username)
	if err != nil && existingUser != nil {
		return errors.New("id already exists")
	}

	return nil
}
