package server

import (
	"fmt"
	"net/http"

	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/go-rendezvous/rendezvous/internal/router"
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

		router.InitRouter(server)

		for _, route := range server.Routes() {
			fmt.Printf("  Method: %-7s Path: %s\n", route.Method, route.Path)
		}

		if err := server.Start(":8080"); err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
		return nil
	},
}
