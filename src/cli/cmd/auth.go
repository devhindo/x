package cmd

import (
	"github.com/devhindo/x/src/cli/auth"
	"github.com/devhindo/x/src/cli/clear"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with X",
	Run: func(cmd *cobra.Command, args []string) {
		verify, _ := cmd.Flags().GetBool("verify")
		clearFlag, _ := cmd.Flags().GetBool("clear")
		url, _ := cmd.Flags().GetBool("url")

		if verify {
			auth.Verify()
		} else if clearFlag {
			clear.StartOver()
		} else if url {
			auth.Get_url_db()
		} else {
			auth.Auth()
		}
	},
}

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.Flags().BoolP("verify", "v", false, "Verify authentication")
	authCmd.Flags().BoolP("clear", "c", false, "Clear authentication")
	authCmd.Flags().Bool("url", false, "Get authorization URL")
}
