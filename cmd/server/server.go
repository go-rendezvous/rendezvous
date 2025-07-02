package server

import (
	"github.com/go-rendezvous/rendezvous/internal/config"
	"github.com/spf13/cobra"
)

var configYml string

func init() {
	ServerCommand.PersistentFlags().StringVarP(&configYml, "config", "c", "settings.yml", "Start server with provided configuration file")
}

var ServerCommand = &cobra.Command{
	Use: "",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.Setup(configYml)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
