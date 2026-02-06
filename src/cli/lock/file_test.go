package lock

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLicenseFileOperations(t *testing.T) {
	// Create a temporary directory to act as HOME
	tempDir, err := os.MkdirTemp("", "locktest")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Set HOME environment variable to tempDir
	originalHome := os.Getenv("HOME")
	defer os.Setenv("HOME", originalHome)
	os.Setenv("HOME", tempDir)

	licenseKey := "test-license-key-123"

	// Test WriteLicenseKeyToFile
	t.Run("WriteLicenseKeyToFile", func(t *testing.T) {
		err := WriteLicenseKeyToFile(licenseKey)
		if err != nil {
			t.Errorf("WriteLicenseKeyToFile() error = %v", err)
		}

		// Verify file exists
		expectedPath := filepath.Join(tempDir, ".tempxcli")
		content, err := os.ReadFile(expectedPath)
		if err != nil {
			t.Errorf("Failed to read license file: %v", err)
		}
		if string(content) != licenseKey {
			t.Errorf("File content = %q, want %q", string(content), licenseKey)
		}
	})

	// Test ReadLicenseKeyFromFile
	t.Run("ReadLicenseKeyFromFile", func(t *testing.T) {
		readKey, err := ReadLicenseKeyFromFile()
		if err != nil {
			t.Errorf("ReadLicenseKeyFromFile() error = %v", err)
		}
		if readKey != licenseKey {
			t.Errorf("ReadLicenseKeyFromFile() = %q, want %q", readKey, licenseKey)
		}
	})

	// Test ClearLicenseFile
	t.Run("ClearLicenseFile", func(t *testing.T) {
		err := ClearLicenseFile()
		if err != nil {
			t.Errorf("ClearLicenseFile() error = %v", err)
		}

		// Verify file is gone
		_, err = ReadLicenseKeyFromFile()
		if err == nil {
			t.Error("ReadLicenseKeyFromFile() should fail after clear, but it succeeded")
		}
		
		expectedPath := filepath.Join(tempDir, ".tempxcli")
		if _, err := os.Stat(expectedPath); !os.IsNotExist(err) {
			t.Error("License file should not exist after clear")
		}
	})
}
