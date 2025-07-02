package service

import (
	"errors"

	"github.com/go-rendezvous/rendezvous/internal/pkg/model"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type Service struct {
	Ctx echo.Context
}

func (s *Service) ExtractClaims() (*model.Claims, error) {
	token, ok := s.Ctx.Get("user").(*jwt.Token) // by default token is stored under `user` key
	if !ok {
		return nil, errors.New("JWT token missing or invalid")
	}
	claimMap, ok := token.Claims.(jwt.MapClaims) // by default claims is of type `jwt.MapClaims`
	if !ok {
		return nil, errors.New("failed to cast claims as jwt.MapClaims")
	}
	return model.ConvertToClaims(claimMap)
}
