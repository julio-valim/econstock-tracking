package domain

import "time"

// StockPrice represents the stock data with relevant fields.
type StockPrice struct {
    Ticker   string    // The stock ticker symbol
    Price    float64   // The current price of the stock
    DateTime time.Time // The timestamp of when the price was recorded
}