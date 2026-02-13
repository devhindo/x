package cmd

import (
	"github.com/devhindo/x/src/cli/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with X",
	Run: func(cmd *cobra.Command, args []string) {
		auth.Auth()
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
}
