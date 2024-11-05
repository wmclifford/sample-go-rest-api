package config_test

import (
	"log"
	"os"
	"testing"

	"go-rest-api/config"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// createMockConfigFile creates a temporary mock config file for testing.
func createMockConfigFile() (string, error) {
	tmpFile, err := os.CreateTemp("", "test_config_*.yaml")
	if err != nil {
		return "", err
	}
	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Fatalf("Could not get current working directory: %v", err)
	// }
	//
	// configDir := filepath.Join(wd, "config")
	// err = os.MkdirAll(configDir, os.ModePerm)
	// if err != nil {
	// 	log.Fatalf("Could not create config directory: %v", err)
	// }
	//
	// filePath := filepath.Join(configDir, "test_config.yaml")
	// file, err := os.Create(filePath)
	// if err != nil {
	// 	log.Fatalf("Could not create mock config file: %v", err)
	// }
	// defer func(file *os.File) {
	// 	err := file.Close()
	// 	if err != nil {
	// 		log.Printf("Could not close mock config file: %v", err)
	// 	}
	// }(file)

	configContent := `
database:
  host: "localhost"
  port: 5432
  user: "test_user"
  dbname: "test_db"
  password: "test_password"
`
	_, err = tmpFile.WriteString(configContent)
	if err != nil {
		log.Fatalf("Could not write to mock config file: %v", err)
	}

	if err := tmpFile.Close(); err != nil {
		return "", err
	}

	return tmpFile.Name(), nil
}

func removeMockConfigFile(filePath string) {
	err := os.Remove(filePath)
	if err != nil {
		log.Fatalf("Could not remove mock config file: %v", err)
	}
}

func TestLoad(t *testing.T) {
	// Print the current working directory
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Could not get current working directory: %v", err)
	}
	log.Printf("Current working directory: %s\n", wd)

	// Setup
	filePath, err := createMockConfigFile()
	if err != nil {
		t.Fatalf("Could not create mock config file: %v", err)
	}
	defer removeMockConfigFile(filePath)

	// Initialize Viper
	v := viper.New()
	v.SetConfigFile(filePath)

	// Assign viper instance to the config
	config.Config = v

	// Test: attempt to load the configuration file
	log.Println("Calling config.Load")
	err = config.Load(filePath)
	if err != nil {
		t.Fatalf("Could not load config file: %v", err)
	}

	// Debug: Confirm the config file used
	log.Printf("Using config file: %s\n", v.ConfigFileUsed())

	assert.Equal(t, "localhost", config.Config.GetString("database.host"))
	assert.Equal(t, 5432, config.Config.GetInt("database.port"))
	assert.Equal(t, "test_user", config.Config.GetString("database.user"))
	assert.Equal(t, "test_db", config.Config.GetString("database.dbname"))
	assert.Equal(t, "test_password", config.Config.GetString("database.password"))
}
