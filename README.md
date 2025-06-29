# econstock-tracking

## Overview
The econstock-tracking project is an autonomous agent designed to monitor stock prices from APIs. It fetches real-time stock prices, analyzes trends, and optionally utilizes a local language model (LLM) to provide insights on stock movements.

## Features
- Fetches current stock prices from the Alpha Vantage API (or another setted API).
- Stores historical stock prices in a PostgreSQL database.
- Performs technical analysis to identify upward trends.
- Optionally integrates with a local LLM for enhanced analysis and commentary.
- Scheduled to run every minute using a cron job.
- (Optional) Exposes a RESTful API for querying stock trends and historical data.

## Project Structure
```
econstock-tracking
├── cmd
│   └── watcher
│       └── main.go            # Entry point of the application
├── config
│   └── config.go             # Configuration loading
├── internal
│   ├── domain
│   │   └── stock.go          # Stock data structures
│   ├── usecase
│   │   └── monitor.go        # Business logic for monitoring stocks
│   ├── infrastructure
│   │   ├── db
│   │   │   └── postgres.go    # Database connection and operations
│   │   ├── api
│   │   │   └── alphavantage.go # API client for fetching stock prices
│   │   └── llm
│   │       └── llama.go       # Integration with local LLM
│   └── delivery
│       ├── cron
│       │   └── cron.go        # Cron job setup
│       └── http
│           └── server.go      # (Optional) HTTP server for REST API
├── pkg
│   └── analyzer
│       └── grafico.go         # Technical analysis functions
├── scripts
│   └── migrations.sql         # SQL scripts for database setup
├── .env.example               # Environment variable template
├── go.mod                     # Module definition and dependencies
└── README.md                  # Project documentation
```

## Setup Instructions
1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd econstock-tracking
   ```

2. **Install dependencies:**
   Ensure you have Go installed, then run:
   ```
   go mod tidy
   ```

3. **Configure environment variables:**
   Copy `.env.example` to `.env` and fill in the required values:
   ```
   cp .env.example .env
   ```

4. **Run database migrations:**
   Execute the SQL scripts in `scripts/migrations.sql` to set up the database tables.

5. **Start the application:**
   Run the application using:
   ```
   go run cmd/watcher/main.go
   ```

## Usage
- The application will automatically fetch stock prices every minute and analyze trends.
- If integrated with the optional HTTP server, you can access endpoints to query stock trends and historical data.

## Future Enhancements
- Implement web scraping for news related to stocks.
- Add machine learning capabilities for predictive analysis.
- Create dashboards for visualizing stock trends.
- Implement notification systems via Telegram, Slack, or email.

## License
This project is licensed under the MIT License. See the LICENSE file for details.