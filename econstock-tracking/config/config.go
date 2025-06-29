package config

import (
    "log"
    "os"
    "strings"
    "github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
    DatabaseURL     string
    APIKey          string
    LLMModelPath    string
    Tickers         []string
    StockAPIProvider string // Adicione este campo
}

// Load reads the configuration from the .env file
func Load() *Config {
    // Load environment variables from .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Create a new Config instance
    return &Config{
        DatabaseURL:      os.Getenv("DB_URL"),
        APIKey:           os.Getenv("API_KEY_ALPHA"),
        LLMModelPath:     os.Getenv("LLM_MODEL_PATH"),
        Tickers:          parseTickers(os.Getenv("TICKERS")),
        StockAPIProvider: os.Getenv("STOCK_API_PROVIDER"), // Novo campo
    }
}

// parseTickers splits the TICKERS string into a slice
func parseTickers(tickers string) []string {
    if tickers == "" {
        return []string{}
    }
    return strings.Split(tickers, ",")
}