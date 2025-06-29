package api

import (
    // "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

type YahooFinanceAPI struct{}

func NewYahooFinanceAPI() *YahooFinanceAPI {
    return &YahooFinanceAPI{}
}

func (y *YahooFinanceAPI) FetchPrice(ticker string) (float64, error) {
    url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=1m&range=1d", ticker)
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return 0, err
    }
    req.Header.Set("User-Agent", "Mozilla/5.0")
    req.Header.Set("Accept", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, err
    }
    defer resp.Body.Close()

    body, _ := io.ReadAll(resp.Body)
    if resp.StatusCode != 200 {
        return 0, fmt.Errorf("Yahoo API error: %s", string(body))
    }

    var data struct {
        Chart struct {
            Result []struct {
                Meta struct {
                    RegularMarketPrice float64 `json:"regularMarketPrice"`
                } `json:"meta"`
            } `json:"result"`
            Error interface{} `json:"error"`
        } `json:"chart"`
    }

    if err := json.Unmarshal(body, &data); err != nil {
        return 0, err
    }

    if len(data.Chart.Result) == 0 {
        return 0, fmt.Errorf("no data for ticker %s", ticker)
    }

    return data.Chart.Result[0].Meta.RegularMarketPrice, nil
}

// //FetchPrice implements StockAPI interface
// func (y *YahooFinanceAPI) FetchPrice(ticker string) (float64, error) {
//     url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?interval=1m&range=1d", ticker)
//     resp, err := http.Get(url)
//     if err != nil {
//         return 0, err
//     }
//     defer resp.Body.Close()

//     body, _ := io.ReadAll(resp.Body)
//     if resp.StatusCode != 200 {
//         return 0, fmt.Errorf("Yahoo API error: %s", string(body))
//     }

//     // Debug: log a resposta bruta (remova em produção)
//     fmt.Println("Yahoo response:", string(body))

//     // Volte o ponteiro do body para o início para o json.NewDecoder funcionar
//     resp.Body = io.NopCloser(bytes.NewBuffer(body))

//     var data struct {
//         Chart struct {
//             Result []struct {
//                 Meta struct {
//                     RegularMarketPrice float64 `json:"regularMarketPrice"`
//                 } `json:"meta"`
//             } `json:"result"`
//             Error interface{} `json:"error"`
//         } `json:"chart"`
//     }

//     if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
//         return 0, err
//     }

//     if len(data.Chart.Result) == 0 {
//         return 0, fmt.Errorf("no data for ticker %s", ticker)
//     }

//     return data.Chart.Result[0].Meta.RegularMarketPrice, nil
// }