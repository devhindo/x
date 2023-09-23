package lock

import (
	"encoding/base64"
	"fmt"
	"github.com/google/uuid"
	"io"
	"os"
	"path/filepath"
)

func GenerateLicenseKey() (string, error) {
	uuid := uuid.New()
	uuidBytes := uuid[:]
	licenseKeyBytes := append(uuidBytes)
	licenseKey := base64.StdEncoding.EncodeToString(licenseKeyBytes)
	return licenseKey, nil
}

func WriteLicenseKeyToFile(licenseKey string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
			return fmt.Errorf("Error getting user home directory: %v", err)
	}

	licenseFilePath := filepath.Join(homeDir, ".tempxcli")
	licenseFile, err := os.Create(licenseFilePath)
	if err != nil {
			return fmt.Errorf("Error creating license file: %v", err)
	}
	defer licenseFile.Close()

	_, err = licenseFile.WriteString(licenseKey)
	if err != nil {
			return fmt.Errorf("Error writing license key to file: %v", err)
	}

	return nil
}

func ReadLicenseKeyFromFile() (string, error) {
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

func ClearLicenseFile() error {
    homeDir, err := os.UserHomeDir()
    if err != nil {
            return fmt.Errorf("Error getting user home directory: %v", err)
    }

    licenseFilePath := filepath.Join(homeDir, ".tempxcli")
    err = os.Remove(licenseFilePath)
    if err != nil {
            return fmt.Errorf("Error deleting license file: %v", err)
    }

    return nil
}