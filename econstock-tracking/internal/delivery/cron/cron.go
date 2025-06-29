package cron

import (
    "log"

    "github.com/robfig/cron/v3"
    "econstock-tracking/internal/usecase"
)

// CronJob struct holds the use case for monitoring stocks
type CronJob struct {
    UseCase *usecase.MonitorUseCase
    Tickers []string
}

// NewCron creates a new CronJob instance
func NewCron(useCase *usecase.MonitorUseCase, tickers []string) *CronJob {
    return &CronJob{
        UseCase: useCase,
        Tickers: tickers,
    }
}

// Start initializes and starts the cron job to run every minute
func (c *CronJob) Start() {
    // Create a new cron scheduler
    scheduler := cron.New()

    // Schedule the monitoring job to run every minute
    _, err := scheduler.AddFunc("@every 1m", func() {
        if err := c.UseCase.Run(c.Tickers); err != nil {
            log.Printf("Error running monitoring use case: %v", err)
        }
    })
    if err != nil {
        log.Fatalf("Error scheduling cron job: %v", err)
    }

    // Start the cron scheduler
    scheduler.Start()
    log.Println("Cron job started, running every minute.")
}