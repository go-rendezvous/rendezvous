package server

import (
	"net/http"
	"time"

	"github.com/go-rendezvous/rendezvous/internal/api"
	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var configYml string

func init() {
	ServerCommand.PersistentFlags().StringVarP(&configYml, "config", "c", "config.yml", "Start server with provided configuration file")
}

var ServerCommand = &cobra.Command{
	Use: "server",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.Setup(configYml)
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		server := echo.New()
		server.Use(middleware.Logger())
		// server.Use(middleware.Recover())

		var id int = 1
		server.GET("/get", func(c echo.Context) error {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"user_id":  id,
				"username": "ethan",
				"exp":      time.Now().Add(time.Hour * 72).Unix(),
			})

			// Sign the token
			logrus.Info(config.GetConf())
			tokenString, err := token.SignedString([]byte(config.GetConf().Secret))
			if err != nil {
				logrus.Error(err)
				return c.JSON(http.StatusBadRequest, err)
			}
			return c.JSON(http.StatusOK, tokenString)
		})

		group := server.Group("/pro", echojwt.WithConfig(echojwt.Config{
			SigningKey: []byte(config.GetConf().Secret),
		}))
		meeting := api.Meeting{}

		group.GET("", meeting.List)
		group.POST("", meeting.Insert)

		if err := server.Start(":8080"); err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
		return nil
	},
}
