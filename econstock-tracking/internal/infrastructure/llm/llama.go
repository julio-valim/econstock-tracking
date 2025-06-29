package llm

import (
    "encoding/json"
    "fmt"
    "os/exec"
)

// LLMClient is a struct that wraps the integration with a local LLM.
type LLMClient struct {
    modelPath string
}

// NewLLMClient initializes a new LLMClient with the given model path.
func NewLLMClient(modelPath string) *LLMClient {
    return &LLMClient{modelPath: modelPath}
}

// AnalyzePattern takes a ticker and historical prices, and returns an analysis of the stock pattern.
func (c *LLMClient) AnalyzePattern(ticker string, historicalPrices []float64) (string, error) {
    input := map[string]interface{}{
        "ticker":    ticker,
        "historico": historicalPrices,
    }

    inputJSON, err := json.Marshal(input)
    if err != nil {
        return "", fmt.Errorf("failed to marshal input: %w", err)
    }

    // Call the LLM model using a command line interface
    cmd := exec.Command(c.modelPath, string(inputJSON))
    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("failed to execute LLM command: %w", err)
    }

    return string(output), nil
}