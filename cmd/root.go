package cmd

import (
    "github.com/spf13/cobra"
    "marketpulse/logger"
)

var rootCmd = &cobra.Command{
    Use:   "marketpulse",
    Short: "MarketPulse CLI provides real-time financial data, news, and sentiment analysis.",
    Long:  `MarketPulse CLI is a tool to aggregate market data, fetch stock information, and analyze financial news.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to be called once to run the root command.
func Execute() error {
    return rootCmd.Execute()
}

func init() {
    // Define flags and configuration settings
    rootCmd.PersistentFlags().StringP("api-key", "a", "", "API key for accessing financial data")
    // rootCmd.MarkPersistentFlagRequired("api-key")
    
    // Log level flag
    rootCmd.PersistentFlags().StringP("log-level", "l", "info", "Log level (debug, info, warn, error)")
    
    // Set log level based on the flag value
    logLevel, _ := rootCmd.PersistentFlags().GetString("log-level")
    logger.SetLogLevel(logLevel) }

