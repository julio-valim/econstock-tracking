package db

// StockRepo define as operações necessárias para o repositório de ações.
type StockRepo interface {
    Save(ticker string, price float64) error
    GetHistory(ticker string, n int) ([]float64, error)
    SaveTrend(ticker string, comment string) error
}