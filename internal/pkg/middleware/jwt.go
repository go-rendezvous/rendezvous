package middleware

import (
	"github.com/go-rendezvous/rendezvous/internal/config"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func JsonWebToken() echo.MiddlewareFunc {
	config := echojwt.Config{
		SigningKey: []byte(config.GetConf().Secret),
	}
	return echojwt.WithConfig(config)
}
