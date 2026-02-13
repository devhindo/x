package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"github.com/devhindo/x/src/cli/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and manage Twitter Apps",
	Long:  `Manage your Twitter Developer Apps locally. Add, remove, and switch between multiple apps.`,
}

var initAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Register a new Twitter App",
	Run: func(cmd *cobra.Command, args []string) {
		var (
			appName      string
			clientID     string
			clientSecret string
		)

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewInput().
					Title("App Name").
					Description("Give your app a unique nickname (e.g. MyBot)").
					Value(&appName).
					Validate(func(s string) error {
						if len(s) == 0 {
							return fmt.Errorf("app name cannot be empty")
						}
						return nil
					}),
				huh.NewInput().
					Title("Client ID").
					Description("From Twitter Developer Portal").
					Value(&clientID).
					Validate(func(s string) error {
						if len(s) == 0 {
							return fmt.Errorf("client ID cannot be empty")
						}
						return nil
					}),
				huh.NewInput().
					Title("Client Secret").
					Description("From Twitter Developer Portal").
					Value(&clientSecret).
					Password(true).
					Validate(func(s string) error {
						if len(s) == 0 {
							return fmt.Errorf("client secret cannot be empty")
						}
						return nil
					}),
			),
		)

		err := form.Run()
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		action := func() {
			cfg, err := config.LoadConfig()
			if err != nil {
				cfg = &config.Config{}
			}

			cfg.AddApp(config.App{
				Name:         appName,
				ClientID:     clientID,
				ClientSecret: clientSecret,
			})

			if cfg.ActiveApp == "" {
				cfg.ActiveApp = appName
			}

			err = config.SaveConfig(cfg)
			if err != nil {
				fmt.Println("Error saving config:", err)
				os.Exit(1)
			}
		}

		_ = spinner.New().Title("Saving app details...").Action(action).Run()

		fmt.Printf("App '%s' added successfully! 🎉\n", appName)
		fmt.Println("Run 'x auth' to authenticate this app.")
	},
}

var initUseCmd = &cobra.Command{
	Use:   "use",
	Short: "Select the active Twitter App",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil || len(cfg.Apps) == 0 {
			fmt.Println("No apps found. Run 'x init add' to register an app.")
			return
		}

		var selectedApp string
		var options []huh.Option[string]

		for _, app := range cfg.Apps {
			label := app.Name
			if app.Name == cfg.ActiveApp {
				label += " (Active)"
			}
			options = append(options, huh.NewOption(label, app.Name))
		}

		form := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select Active App").
					Options(options...).
					Value(&selectedApp),
			),
		)

		err = form.Run()
		if err != nil {
			fmt.Println("Operation cancelled.")
			return
		}

		cfg.SetActiveApp(selectedApp)
		err = config.SaveConfig(cfg)
		if err != nil {
			fmt.Println("Error saving config:", err)
			os.Exit(1)
		}

		fmt.Printf("Switched to '%s' successfully! 🚀\n", selectedApp)
	},
}

var initDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a registered Twitter App",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfig()
		if err != nil || len(cfg.Apps) == 0 {
			fmt.Println("No apps found to delete.")
			return
		}

		var appToDelete string
		var options []huh.Option[string]

		for _, app := range cfg.Apps {
			options = append(options, huh.NewOption(app.Name, app.Name))
		}

		// Step 1: Select App
		selectForm := huh.NewForm(
			huh.NewGroup(
				huh.NewSelect[string]().
					Title("Select App to Delete").
					Options(options...).
					Value(&appToDelete),
			),
		)

		err = selectForm.Run()
		if err != nil {
			fmt.Println("Operation cancelled.")
			return
		}

		// Step 2: Confirm
		var confirm bool
		confirmForm := huh.NewForm(
			huh.NewGroup(
				huh.NewConfirm().
					Title(fmt.Sprintf("Are you sure you want to delete '%s'?", appToDelete)).
					Affirmative("Yes, delete it").
					Negative("No, keep it").
					Value(&confirm),
			),
		)

		err = confirmForm.Run()
		if err != nil {
			fmt.Println("Operation cancelled.")
			return
		}

		if confirm {
			cfg.RemoveApp(appToDelete)
			err = config.SaveConfig(cfg)
			if err != nil {
				fmt.Println("Error saving config:", err)
				os.Exit(1)
			}
			fmt.Printf("App '%s' deleted. 🗑️\n", appToDelete)
		} else {
			fmt.Println("Deletion cancelled.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.AddCommand(initAddCmd)
	initCmd.AddCommand(initUseCmd)
	initCmd.AddCommand(initDeleteCmd)
}
