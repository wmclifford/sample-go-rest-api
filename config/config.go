package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

// Config is a global variable that holds the configuration data.
var Config *viper.Viper

// Load initializes the Viper configuration library and loads the
// configuration file specified.
func Load(file string) error {
	if Config == nil {
		Config = viper.New()
	}

	// Set up Viper to read from the config file
	Config.SetConfigFile(file)
	Config.SetConfigType("yaml")

	// Read configuration from the file
	if err := Config.ReadInConfig(); err != nil {
		log.Printf("Could not load config file: %v", err)
	}

	// Set environment variable prefix and automatic override
	Config.SetEnvPrefix("APP")
	Config.AutomaticEnv()
	Config.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	return nil
}
