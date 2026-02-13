package clear

import (
	"fmt"
	"os"

	"github.com/devhindo/x/src/cli/config"
)

func StartOver() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	app := cfg.GetActiveApp()
	if app == nil {
		fmt.Println("No active app found.")
		return
	}

	app.User = nil
	cfg.AddApp(*app) // Update app

	err = config.SaveConfig(cfg)
	if err != nil {
		fmt.Println("Error clearing user:", err)
		os.Exit(1)
	}

	fmt.Println("User session cleared successfully.")
}
