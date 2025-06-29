package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// StockAPI defines the interface for fetching stock prices.
type StockAPI interface {
	FetchPrice(ticker string) (float64, error)
}

// AlphaVantageAPI implements the StockAPI interface for Alpha Vantage.
type AlphaVantageAPI struct {
	APIKey string
}

// NewAlphaVantage creates a new instance of AlphaVantageAPI.
func NewAlphaVantage(apiKey string) *AlphaVantageAPI {
	return &AlphaVantageAPI{APIKey: apiKey}
}

// FetchPrice fetches the current stock price for the given ticker from Alpha Vantage.
func (a *AlphaVantageAPI) FetchPrice(ticker string) (float64, error) {
	url := fmt.Sprintf("https://www.alphavantage.co/query?function=TIME_SERIES_INTRADAY&symbol=%s&interval=1min&apikey=%s", ticker, a.APIKey)

	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch price: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("error response from API: %s", resp.Status)
	}

	var result struct {
		TimeSeries map[string]struct {
			Close string `json:"4. close"`
		} `json:"Time Series (1min)"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("failed to decode response: %w", err)
	}

	// Get the latest time series entry
	for _, v := range result.TimeSeries {
		price, err := strconv.ParseFloat(v.Close, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse price: %w", err)
		}
		return price, nil
	}

	return 0, fmt.Errorf("no price data available for ticker: %s", ticker)
}
