package auth

import (
	"os"
	"fmt"
	"io"
	"path/filepath"
)

func (u *User) RetrieveLicenseKey() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
			return "", fmt.Errorf("Error getting user home directory: %v", err)
	}

	licenseFilePath := filepath.Join(homeDir, ".tempxcli")
	licenseFile, err := os.Open(licenseFilePath)
	if err != nil {
			return "", fmt.Errorf("Error opening license file: %v", err)
	}
	defer licenseFile.Close()

	licenseFileBytes, err := io.ReadAll(licenseFile)
	if err != nil {
			return "", fmt.Errorf("Error reading license file: %v", err)
	}

	licenseKey := string(licenseFileBytes)
	return licenseKey, nil
}