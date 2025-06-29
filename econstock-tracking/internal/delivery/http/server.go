package http

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "econstock-tracking/internal/usecase"
)

// Server represents the HTTP server
type Server struct {
    router      *mux.Router
    monitorUseCase *usecase.MonitorUseCase
}

// NewServer creates a new HTTP server
func NewServer(monitorUseCase *usecase.MonitorUseCase) *Server {
    s := &Server{
        router:      mux.NewRouter(),
        monitorUseCase: monitorUseCase,
    }
    s.routes()
    return s
}

// routes sets up the HTTP routes
func (s *Server) routes() {
    s.router.HandleFunc("/acoes-em-tendencia", s.getAcoesEmTendencia).Methods("GET")
    s.router.HandleFunc("/historico/{ticker}", s.getHistorico).Methods("GET")
    s.router.HandleFunc("/reavaliar/{ticker}", s.reavaliar).Methods("POST")
}

// Start runs the HTTP server
func (s *Server) Start(addr string) error {
    return http.ListenAndServe(addr, s.router)
}

// getAcoesEmTendencia handles requests to get trending stocks
func (s *Server) getAcoesEmTendencia(w http.ResponseWriter, r *http.Request) {
    // Logic to retrieve trending stocks
    // This is a placeholder for actual implementation
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Ações em tendência")
}

// getHistorico handles requests to get historical data for a specific ticker
func (s *Server) getHistorico(w http.ResponseWriter, r *http.Request) {
    ticker := mux.Vars(r)["ticker"]
    // Logic to retrieve historical data for the ticker
    // This is a placeholder for actual implementation
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Histórico para " + ticker)
}

// reavaliar handles requests to re-evaluate a specific ticker
func (s *Server) reavaliar(w http.ResponseWriter, r *http.Request) {
    ticker := mux.Vars(r)["ticker"]
    // Logic to trigger re-evaluation for the ticker
    // This is a placeholder for actual implementation
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Reavaliação para " + ticker)
}