package cmd

import (
    "fmt"
    "os"
    "text/tabwriter"
    "sort"
    "github.com/spf13/cobra"
    "marketpulse/api"
    "marketpulse/logger"
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
    Use:   "fetch [symbol]",
    Short: "Fetches market data for a specified symbol",
    Long:  `Fetches real-time financial data for the given stock symbol from Alpha Vantage and prints the results.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        symbol := args[0]

        // Load configuration (can come from env, file, or defaults)
        if err := api.LoadConfig(); err != nil {
            logger.Error(fmt.Sprintf("Error loading configuration: %v", err))
            return
        }

        apiKey, _ := cmd.Flags().GetString("api-key")

        // Get the API key from config (environment or config file)
        if apiKey == "" {
            apiKey = api.GetAPIKey()
            if apiKey == "" {
                logger.Error("API key is required. Please set it in the configuration or environment variables.")
                return
            }
        }

        // Fetch data
        data, err := api.FetchStockData(symbol, apiKey)
        if err != nil {
            logger.Error(fmt.Sprintf("Error fetching data: %v", err))
            return
        }

        logger.Info(fmt.Sprintf("Market data for %s:", symbol))
        displayData(data)
    },
}

// displayData formats and prints the market data in a table
func displayData(data interface{}) {
    // Ensure the data is a map
    dataMap, ok := data.(map[string]interface{})
    if !ok {
        logger.Error("Unexpected data format, expected a map.")
        return
    }

    // Extract and format the "Time Series (5min)" data
    timeSeriesRaw, ok := dataMap["Time Series (5min)"]
    if !ok {
        logger.Error("Time Series data not found.")
        return
    }

    // Ensure the time series is a map
    timeSeries, ok := timeSeriesRaw.(map[string]interface{})
    if !ok {
        logger.Error("Unexpected format for Time Series data.")
        return
    }

    // Create a tab writer for clean table formatting
    w := new(tabwriter.Writer)
    w.Init(os.Stdout, 0, 8, 2, ' ', 0)

    // Print header
    fmt.Fprintln(w, "Time\tOpen\tHigh\tLow\tClose\tVolume")
    fmt.Fprintln(w, "----\t----\t----\t---\t-----\t------")

    // Get the keys (timestamps) and sort them to display data in chronological order
    timestamps := make([]string, 0, len(timeSeries))
    for timestamp := range timeSeries {
        timestamps = append(timestamps, timestamp)
    }
    sort.Strings(timestamps)

    // Loop over sorted timestamps and extract the corresponding market data
    for _, timestamp := range timestamps {
        recordRaw, ok := timeSeries[timestamp]
        if !ok {
            logger.Error(fmt.Sprintf("No data found for timestamp %s", timestamp))
            continue
        }

        // Ensure record is a map
        record, ok := recordRaw.(map[string]interface{})
        if !ok {
            logger.Error(fmt.Sprintf("Unexpected data format for record at %s", timestamp))
            continue
        }

        open := record["1. open"]
        high := record["2. high"]
        low := record["3. low"]
        close := record["4. close"]
        volume := record["5. volume"]

        // Print formatted data
        fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\n", timestamp, open, high, low, close, volume)
    }

    // Flush the tab writer to output the formatted data
    w.Flush()
}

func init() {
    rootCmd.AddCommand(fetchCmd)
}

