package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "marketpulse",
    Short: "MarketPulse CLI provides real-time financial data, news, and sentiment analysis.",
    Long:  `MarketPulse CLI is a tool to aggregate market data, fetch stock information, and analyze financial news.`,
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("Use 'marketpulse fetch' to get market data.")
    },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to be called once to run the root command.
func Execute() error {
    return rootCmd.Execute()
}

func init() {
    // Define flags and configuration settings
    rootCmd.PersistentFlags().StringP("api-key", "a", "", "API key for accessing financial data")
    rootCmd.MarkPersistentFlagRequired("api-key")
}
