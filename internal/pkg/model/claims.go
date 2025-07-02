package model

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId   int
	Username string
}

func ConvertToClaims(claims jwt.MapClaims) (*Claims, error) {
	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("cannot convert to claims")
	}
	userId, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("cannot convert to claims")
	}
	return &Claims{
		UserId:   int(userId),
		Username: username,
	}, nil
}
