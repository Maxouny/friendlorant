package models

import "time"

type User struct {
	ID          uint      `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"-"`
	Email       string    `json:"email"`
	Token       string    `json:"token"`
	Image       string    `json:"image"`
	ValorantID  int       `json:"valorant_id"`
	UserRating  int       `json:"user_rating"`
	TokenExpire time.Time `json:"token_expire"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
