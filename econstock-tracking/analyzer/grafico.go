package analyzer


// GraficoAnalyzer is responsible for performing technical analysis on stock price data.
type GraficoAnalyzer struct{}

// IsUpTrend checks if the stock price history indicates an upward trend.
// It compares the short-term moving average (7 periods) with the long-term moving average (21 periods)
// and checks if the price history is in a rising sequence.
func (a *GraficoAnalyzer) IsUpTrend(historico []float64) bool {
    if len(historico) < 21 {
        return false // Not enough data to analyze
    }

    mm7 := calcularMedia(historico[len(historico)-7:])  // Short-term moving average
    mm21 := calcularMedia(historico[len(historico)-21:]) // Long-term moving average

    return mm7 > mm21 && crescente(historico) // Check if short-term average is greater and if the prices are rising
}

// calcularMedia calculates the average of a slice of float64 values.
func calcularMedia(precos []float64) float64 {
    total := 0.0
    for _, preco := range precos {
        total += preco
    }
    return total / float64(len(precos))
}

// crescente checks if the price history is in a rising sequence.
func crescente(historico []float64) bool {
    for i := 1; i < len(historico); i++ {
        if historico[i] <= historico[i-1] {
            return false // Found a price that is not greater than the previous one
        }
    }
    return true // All prices are in a rising sequence
}

// NewGraficoAnalyzer cria uma nova instância de GraficoAnalyzer.
func NewGraficoAnalyzer() *GraficoAnalyzer {
    return &GraficoAnalyzer{}
}