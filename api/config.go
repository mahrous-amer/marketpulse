package api

import (
    "github.com/spf13/viper"
    "marketpulse/logger"
)

// LoadConfig loads the configuration using viper.
func LoadConfig() error {

    // Default values
    viper.SetDefault("API_KEY", "xxx")

    // Automatically read environment variables
    viper.AutomaticEnv()

    // Load from config file (if needed)
    viper.SetConfigName("config")
    viper.AddConfigPath(".")
    viper.SetConfigType("yaml")

    // Try reading the config file
    if err := viper.ReadInConfig(); err != nil {
        logger.Warn("No config file found, using environment variables")
    }

    return nil
}

// GetAPIKey returns the API key from environment or config.
func GetAPIKey() string {
    return viper.GetString("API_KEY")
}

