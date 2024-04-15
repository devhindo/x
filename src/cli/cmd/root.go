package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	version = "1.1.3"
	man = fmt.Sprintf(`interact with x (twitter) from terminal. version: %s.

	USAGE
	  x <command>
	
	Commands
	  -h             show this help
	  auth           start authorizing your X account
	  auth --url     get auth url if it didn't open browser after running 'x auth'
	  auth -v        verify authorization after running 'x auth'
	  -t "text"      post a tweet
	  -v             show version (%s)
	  -c             clear authorized account`, version, version)
)

var rootCmd = &cobra.Command{
	Use:   "x",
	Short: "x is a CLI tool for posting tweets",
	Long: man,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(man)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

}
