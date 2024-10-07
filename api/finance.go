package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
    "marketpulse/logger"
)

// Alpha Vantage API URL
const alphaVantageURL = "https://www.alphavantage.co/query"

// FetchStockData fetches market data for the given symbol
func FetchStockData(symbol, apiKey string) (interface{}, error) {
    url := fmt.Sprintf("%s?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", alphaVantageURL, symbol, apiKey)
    logger.Info(fmt.Sprintf("Fetching data for symbol %s from Alpha Vantage", symbol))

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var data interface{}
    if err := json.Unmarshal(body, &data); err != nil {
        return nil, err
    }

    logger.Info(fmt.Sprintf("Successfully fetched data for %s", symbol))
    return data, nil
}

