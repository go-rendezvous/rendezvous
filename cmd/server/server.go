package server

import (
	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/go-rendezvous/rendezvous/internal/store/database"
	"github.com/go-rendezvous/rendezvous/pkg/db"
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
		driver, err := db.NewDBDriver(config.GetConf())
		if err != nil {
			return err
		}
		database.Factory(driver.GetOrm())
		return nil
	},
}
