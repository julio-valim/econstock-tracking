package main

import (
    "log"
    "time"

    "econstock-tracking/config"
    "econstock-tracking/internal/delivery/cron"
    "econstock-tracking/internal/infrastructure/api"
    "econstock-tracking/internal/infrastructure/db"
    "econstock-tracking/internal/infrastructure/llm"
    "econstock-tracking/internal/usecase"
    "econstock-tracking/pkg/analyzer"
)

func main() {
    // Load configuration from .env file
    cfg := config.Load()

    // Establish a connection to the PostgreSQL database
    database := db.NewPostgres(cfg.DatabaseURL)
    if database == nil {
        log.Fatal("Failed to connect to the database")
    }

    // Initialize the stock repository
    stockRepo := db.NewStockRepo(database)

    // Create an instance of stock API client
    var stockAPI api.StockAPI
    switch cfg.StockAPIProvider {
    case "alpha":
        stockAPI = api.NewAlphaVantage(cfg.APIKey)
    case "yahoo":
        stockAPI = api.NewYahooFinanceAPI()
    default:
        log.Fatal("Invalid STOCK_API_PROVIDER. Use 'alpha' or 'yahoo'")
    }

    // Initialize the graphical analyzer
    graficoAnalyzer := analyzer.NewGraficoAnalyzer()

    // Set up the LLM client for optional analysis
    llmClient := llm.NewLLMClient(cfg.LLMModelPath)

    // Create the monitoring use case
    monitorUseCase := usecase.NewMonitorUseCase(stockAPI, stockRepo, graficoAnalyzer, llmClient)

    // Set up the cron job to run the monitoring use case every minute
    cronJob := cron.NewCron(monitorUseCase, cfg.Tickers)
    cronJob.Start()

    // Keep the application running
    select {
    case <-time.After(time.Hour * 24): // Run for 24 hours
        log.Println("Application running for 24 hours, shutting down...")
    }
}