package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

type PostgresRepo struct {
    db *sql.DB
}

func NewStockRepo(db *sql.DB) StockRepo {
    return &PostgresRepo{db: db}
}

func (r *PostgresRepo) Save(ticker string, price float64) error {
    query := `
        INSERT INTO precos_acoes (ticker, preco, datahora)
        VALUES ($1, $2, NOW())
    `
    _, err := r.db.Exec(query, ticker, price)
    if err != nil {
        return fmt.Errorf("failed to save stock price: %w", err)
    }
    return nil
}

func (r *PostgresRepo) GetHistory(ticker string, n int) ([]float64, error) {
    query := `
        SELECT preco FROM precos_acoes
        WHERE ticker = $1
        ORDER BY datahora DESC
        LIMIT $2
    `
    rows, err := r.db.Query(query, ticker, n)
    if err != nil {
        return nil, fmt.Errorf("failed to retrieve stock history: %w", err)
    }
    defer rows.Close()

    var prices []float64
    for rows.Next() {
        var price float64
        if err := rows.Scan(&price); err != nil {
            return nil, fmt.Errorf("failed to scan stock price: %w", err)
        }
        prices = append(prices, price)
    }

    return prices, nil
}

func (r *PostgresRepo) SaveTrend(ticker string, comment string) error {
    query := `
        INSERT INTO acoes_em_tendencia (ticker, motivo, datahora)
        VALUES ($1, $2, NOW())
    `
    _, err := r.db.Exec(query, ticker, comment)
    if err != nil {
        return fmt.Errorf("failed to save trend: %w", err)
    }
    return nil
}

func NewPostgres(url string) *sql.DB {
    db, err := sql.Open("postgres", url)
    if err != nil {
        fmt.Printf("Failed to open database: %v\n", err)
        return nil
    }
    if err := db.Ping(); err != nil {
        fmt.Printf("Failed to ping database: %v\n", err)
        return nil
    }
    return db
}