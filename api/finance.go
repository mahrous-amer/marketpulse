package api

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

// Alpha Vantage API URL
const baseURL = "https://www.alphavantage.co/query"

// FetchStockData fetches market data for the given symbol
func FetchStockData(symbol, apiKey string) (interface{}, error) {
    url := fmt.Sprintf("%s?function=TIME_SERIES_INTRADAY&symbol=%s&interval=5min&apikey=%s", baseURL, symbol, apiKey)
    
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

    return data, nil
}

