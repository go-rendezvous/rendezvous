package server

import "github.com/spf13/cobra"

var ServerCommand = &cobra.Command{
	Use: "",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
