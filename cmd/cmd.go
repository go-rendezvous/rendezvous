package cmd

import (
	"github.com/go-rendezvous/rendezvous/cmd/server"
	"github.com/spf13/cobra"
)

func init() {
	RootCommand.AddCommand(server.ServerCommand)
}

var RootCommand = &cobra.Command{
	Use: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
