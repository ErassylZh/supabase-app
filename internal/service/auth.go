package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

type Auth interface {
	VerifyToken(tokenString string) (userId string, err error)
}

type AuthService struct {
	secret string
}

func NewAuthService(secret string) *AuthService {
	return &AuthService{secret: secret}
}

func (a *AuthService) VerifyToken(tokenString string) (userId string, err error) {
	tokenMain := strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(tokenMain, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(a.secret), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if sub, ok := claims["sub"].(string); ok {
			return sub, nil
		}
		return "", fmt.Errorf("user_id in claims empty")
	}
	return "", fmt.Errorf("can't read claims")
}
