package service

import (
	"net/http"
	"time"

	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  1,
		"username": "ethan",
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	// Sign the token
	tokenString, err := token.SignedString([]byte(config.GetConf().Secret))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusOK, tokenString)
}

func Logout(c echo.Context) error {
	return nil
}
