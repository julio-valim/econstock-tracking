package usecase

import (
    "econstock-tracking/internal/infrastructure/api"
    "econstock-tracking/internal/infrastructure/llm"
    "econstock-tracking/pkg/analyzer"
    "econstock-tracking/internal/infrastructure/db"
)

// MonitorUseCase encapsulates the business logic for monitoring stock prices.
type MonitorUseCase struct {
    API      api.StockAPI
    Repo     db.StockRepo
    Analyzer *analyzer.GraficoAnalyzer
    LLM      *llm.LLMClient
}

// Run executa o monitoramento para os tickers fornecidos.
func (u *MonitorUseCase) Run(tickers []string) error {
    for _, ticker := range tickers {
        price, err := u.API.FetchPrice(ticker)
        if err != nil {
            return err
        }
        err = u.Repo.Save(ticker, price)
        if err != nil {
            return err
        }
        history, err := u.Repo.GetHistory(ticker, 21)
        if err != nil {
            return err
        }
        if u.Analyzer.IsUpTrend(history) {
            comment, err := u.LLM.AnalyzePattern(ticker, history)
            if err != nil {
                return err
            }
            err = u.Repo.SaveTrend(ticker, comment)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

// NewMonitorUseCase cria uma nova instância de MonitorUseCase.
func NewMonitorUseCase(
    api api.StockAPI,
    repo db.StockRepo,
    analyzer *analyzer.GraficoAnalyzer,
    llm *llm.LLMClient,
) *MonitorUseCase {
    return &MonitorUseCase{
        API:      api,
        Repo:     repo,
        Analyzer: analyzer,
        LLM:      llm,
    }
}