package cmd

import (
    "fmt"
    "os"
    "reflect"
    "text/tabwriter" // Import tabwriter for table formatting
    "github.com/fatih/color"
    "github.com/spf13/cobra"
    "marketpulse/api" // Import your API package
)

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
    Use:   "fetch",
    Short: "Fetches market data for a specified symbol",
    Long:  `Fetches real-time financial data for the given stock symbol from Alpha Vantage and prints the results.`,
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        symbol := args[0]

        // Get API key from flag
        apiKey, _ := cmd.Flags().GetString("api-key")
        if apiKey == "" {
            // Try to get API key from environment variable if not provided via flag
            apiKey = os.Getenv("ALPHA_VANTAGE_API_KEY")
        }
        if apiKey == "" {
            color.Red("API key is required. Please provide it using the --api-key flag or the ALPHA_VANTAGE_API_KEY environment variable.")
            return
        }

        // Fetch data
        data, err := api.FetchStockData(symbol, apiKey)
        if err != nil {
            color.Red("Error fetching data: %v", err)
            return
        }

        color.Cyan("Market data for %s:", symbol)
        // Display data in a colored format or structured manner
        fmt.Println(data)
                // Use reflection to handle unknown types
        v := reflect.ValueOf(data)
        if v.Kind() != reflect.Map {
            color.Red("Unexpected data format")
            return
        }

        // Create a new tab writer to format the data in a table
        w := new(tabwriter.Writer)
        w.Init(os.Stdout, 0, 8, 2, ' ', 0)

        // Print header
        fmt.Fprintln(w, "Key\tValue")
        fmt.Fprintln(w, "----\t-----")

        // Print data in tabular format
       for _, key := range v.MapKeys() {
            value := v.MapIndex(key)
            fmt.Fprintf(w, "%v\t%v\n", key.Interface(), value.Interface())
        }
        // Flush the tab writer to output the formatted data
        w.Flush()
    },
}

func init() {
    rootCmd.AddCommand(fetchCmd)
}

