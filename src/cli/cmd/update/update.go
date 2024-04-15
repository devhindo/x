package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the CLI to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		cmdUpdate()
	},
}

func cmdUpdate() {
	fmt.Println("Updating CLI...")

	if !isGoInstalled() {
		fmt.Println("")
	}

	cmd := exec.Command("go", "get", "-u", "github.com/bradford-hamilton/monkey")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("CLI updated successfully!")
}

func isGoInstalled() bool {
	_, err := exec.LookPath("go")
	if err != nil {
		return false
	}
	return true
}

func updateUsingGo() error {
	cmd := exec.Command("go", "get", "-u", "github.com/bradford-hamilton/monkey")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("CLI updated successfully!")
	return nil
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

//TODO: x update -v 1.1.1 (default for v:latest)