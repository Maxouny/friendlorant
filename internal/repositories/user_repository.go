package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"friendlorant/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users
	(username, password, email, image, valorant_id, user_rating, token_expire, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	err := ur.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email, user.Image, user.ValorantID, user.UserRating, user.TokenExpire, user.CreatedAt, user.UpdatedAt).
		Scan(&user.ID)
	if err != nil {
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User

	query := `SELECT * FROM users WHERE id = $1`
	err := ur.db.QueryRowContext(ctx, query, id).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Image, &user.ValorantID, &user.UserRating, &user.Token, &user.TokenExpire, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT * FROM users WHERE email = $1`
	err := ur.db.QueryRowContext(ctx, query, email).
		Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Image, &user.ValorantID, &user.UserRating, &user.Token, &user.TokenExpire, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET username = $1, password = $2, email = $3, image = $4, valorant_id = $5, user_rating = $6, token_expire = $7, updated_at = $8 WHERE id = $9`
	_, err := ur.db.ExecContext(ctx, query, user.Username, user.Password, user.Email, user.Image, user.ValorantID, user.UserRating, user.TokenExpire, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := ur.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
