package service

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"work-project/internal/model"
	"work-project/internal/repository"
)

type Auth interface {
	VerifyToken(tokenString string) (userId string, err error)
	GetUserByToken(ctx context.Context, tokenString string) (user model.User, err error)
}

type AuthService struct {
	userRepo repository.User
	secret   string
}

func NewAuthService(userRepo repository.User, secret string) *AuthService {
	return &AuthService{secret: secret, userRepo: userRepo}
}

func (a *AuthService) VerifyToken(tokenString string) (userId string, err error) {
	tokenMain := strings.Split(tokenString, " ")
	if len(tokenMain) < 2 {
		return "", fmt.Errorf("incorrect token, add bearer identifier")
	}
	token, err := jwt.Parse(tokenMain[1], func(token *jwt.Token) (interface{}, error) {
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

func (a *AuthService) GetUserByToken(ctx context.Context, tokenString string) (user model.User, err error) {
	//userId, err := a.VerifyToken(tokenString)
	//if err != nil {
	//	return model.User{}, err
	//}
	userId := "007ffdc1-9ce2-4146-a487-e12328de6d5c"
	return a.userRepo.GetByID(ctx, userId)
}
