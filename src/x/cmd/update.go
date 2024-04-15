package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&vFlag, "version", "v", "latest", "Update the CLI to specific version")
}

var (
	vFlag string

	updateCmd = &cobra.Command{
		Use:   "update",
		Short: "Update the CLI to specific version ",
		Long: `Update the CLI to the latest version or specify a version using -v flag
				Example: x update -v 1.1.1
				default: x update -v latest
		`,
		RunE: update,
	}
)

func update(cmd *cobra.Command, args []string) error {

	fmt.Println("Updating CLI...")

	vFlag, err := cmd.Flags().GetString("version")
	if err != nil {
		err = fmt.Errorf("error getting version flag: %v", err)
		return err
	}

	// validate version
	if vFlag != "latest" {
		err = validateVersion(vFlag)
		if err != nil {
			return err
		}
	}

	if isGoInstalled() {
		err = updateUsingGo()
		if err != nil {
			return err
		}
		return nil
	}

	return nil
}

func validateVersion(v string) error {
	// Regular expression pattern for semantic versioning
	pattern := `^v\d+\.\d+\.\d+$`
	match, err := regexp.MatchString(pattern, v)
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("invalid version format. Please use semantic versioning like v1.1.1")
	}
	return nil
}

func cmdUpdate() {
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

	if isGoInstalled() {
		fmt.Println("Go is installed...")

		return
	}

	fmt.Println("Go is not installed...")

	fmt.Println("CLI updated successfully!")
}

func isGoInstalled() bool {
	_, err := exec.LookPath("go")

	return err == nil
}

func updateUsingGo() error {
	cmd := exec.Command("go", "get", "-u", "github.com/devhindo/x/src/cli/cmd")
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

//TODO: x update -v 1.1.1 (default for v:latest)
