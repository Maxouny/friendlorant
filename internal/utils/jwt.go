package utils

import (
	"fmt"
	"time"

	"friendlorant/internal/config"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID uint) (string, error) {
	envCfg, err := config.LoadEnvConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %v", err)
	}

	claims := &JWTClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
			NotBefore: time.Now().Unix(),
			Issuer:    "friendlorant",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(envCfg.JWT))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	envCfg, err := config.LoadEnvConfig()
	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(envCfg.JWT), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
