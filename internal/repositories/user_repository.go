package repositories

import (
	"context"
	"fmt"
	"strings"

	"friendlorant/internal/models"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, id uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	UpdateUser(ctx context.Context, user *models.User) error
	DeleteUser(ctx context.Context, id uint) error
}

type userRepository struct {
	pgx *pgx.Conn
}

func NewUserRepository(pgx *pgx.Conn) UserRepository {
	return &userRepository{
		pgx: pgx,
	}
}

func (ur *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users
	(username, email, password, token, image, valorant_id, user_rating, token_expire, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err := ur.pgx.QueryRow(ctx, query, user.Username, user.Email, &user.Password, user.Token, user.Image, user.ValorantID, user.UserRating, user.TokenExpire, user.CreatedAt, user.UpdatedAt).
		Scan(&user.ID)
	if err != nil {
		pgErr, ok := err.(*pgconn.PgError)
		if ok && pgErr.Code == "23505" {
			if strings.Contains(pgErr.Detail, "username") {
				return fmt.Errorf("username %s already exists", user.Username)
			} else if strings.Contains(pgErr.Detail, "email") {
				return fmt.Errorf("email %s already exists", user.Email)
			}
		}
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}

func (ur *userRepository) GetUserByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User

	query := `SELECT id, username, email, password, image, valorant_id, user_rating, token, token_expire, created_at, updated_at FROM users WHERE id = $1`
	err := ur.pgx.QueryRow(ctx, query, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Image, &user.ValorantID, &user.UserRating, &user.Token, &user.TokenExpire, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to scan user: %v", err)
	}

	return &user, nil
}

func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, image, valorant_id, user_rating, token, token_expire, created_at, updated_at FROM users WHERE email = $1`
	err := ur.pgx.QueryRow(ctx, query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Image, &user.ValorantID, &user.UserRating, &user.Token, &user.TokenExpire, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password, image, valorant_id, user_rating, token, token_expire, created_at, updated_at FROM users WHERE username = $1`
	err := ur.pgx.QueryRow(ctx, query, username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Image, &user.ValorantID, &user.UserRating, &user.Token, &user.TokenExpire, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *userRepository) UpdateUser(ctx context.Context, user *models.User) error {
	query := `UPDATE users SET username = $1, email = $2, password = $3, image = $4, valorant_id = $5, user_rating = $6, token_expire = $7, updated_at = $8 WHERE id = $9`
	_, err := ur.pgx.Exec(ctx, query, user.Username, user.Email, &user.Password, user.Image, user.ValorantID, user.UserRating, user.TokenExpire, user.UpdatedAt, user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) DeleteUser(ctx context.Context, id uint) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := ur.pgx.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
