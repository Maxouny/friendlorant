package models

import "time"

type User struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username" validate:"required,min=3,max=32"`
	Email       string    `json:"email" validate:"required,min=3,max=32"`
	Password    string    `json:"password,omitempty" validate:"required,min=6"`
	Token       string    `json:"token"`
	Image       string    `json:"image"`
	ValorantID  int       `json:"valorant_id"`
	UserRating  int       `json:"user_rating"`
	TokenExpire time.Time `json:"token_expire"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
type PublickUser struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Image     string    `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
