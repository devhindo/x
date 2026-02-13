package auth

import (
	"fmt"
	"os"

	"github.com/devhindo/x/src/cli/config"
)

func Auth() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	if len(cfg.Apps) == 0 {
		fmt.Println("No apps found. Run 'x init add' first.")
		os.Exit(1)
	}

	app := cfg.GetActiveApp()
	if app == nil {
		// Fallback to first app if active not set
		if len(cfg.Apps) > 0 {
			app = &cfg.Apps[0]
			cfg.ActiveApp = app.Name
			config.SaveConfig(cfg)
		} else {
			fmt.Println("No apps found. Run 'x init add' first.")
			os.Exit(1)
		}
	}

	fmt.Printf("Authenticating app '%s'...\n", app.Name)
	err = StartAuthFlow(app)
	if err != nil {
		fmt.Println("Authentication failed:", err)
		os.Exit(1)
	}

	fmt.Println("Authentication successful! You can now use 'x tweet'.")
}
