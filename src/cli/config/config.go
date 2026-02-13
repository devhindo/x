package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type User struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Expiry       time.Time `json:"expiry"`
}

type App struct {
	Name         string `json:"name"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	User         *User  `json:"user"`
}

type Config struct {
	ActiveApp string `json:"active_app"`
	Apps      []App  `json:"apps"`
}

func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	configDir := filepath.Join(homeDir, ".x-cli")
	if err := os.MkdirAll(configDir, 0700); err != nil {
		return "", err
	}
	return filepath.Join(configDir, "config.json"), nil
}

func LoadConfig() (*Config, error) {
	path, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &Config{Apps: []App{}}, nil
		}
		return nil, err
	}

	var cfg Config
	if err := json.Unmarshal(file, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	path, err := GetConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0600)
}

func (c *Config) AddApp(app App) {
	// Check if app with same name exists, update it if so
	for i, a := range c.Apps {
		if a.Name == app.Name {
			c.Apps[i] = app
			return
		}
	}
	c.Apps = append(c.Apps, app)
}

func (c *Config) RemoveApp(name string) {
	var newApps []App
	for _, app := range c.Apps {
		if app.Name != name {
			newApps = append(newApps, app)
		}
	}
	c.Apps = newApps
	if c.ActiveApp == name {
		c.ActiveApp = ""
	}
}

func (c *Config) SetActiveApp(name string) error {
	for _, app := range c.Apps {
		if app.Name == name {
			c.ActiveApp = name
			return nil
		}
	}
	return fmt.Errorf("app %s not found", name)
}

func (c *Config) GetActiveApp() *App {
	for i, app := range c.Apps {
		if app.Name == c.ActiveApp {
			return &c.Apps[i]
		}
	}
	return nil
}
