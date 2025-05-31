package http

import (
	"encoding/json"
	"net/http"
	"intelligent-calculator/calculator/entity"
	"intelligent-calculator/calculator/endpoint"
	"context"
)

type Handler struct {
	Endpoint *endpoint.CalcEndpoint
}

func NewHandler(ep *endpoint.CalcEndpoint) *Handler {
	return &Handler{Endpoint: ep}
}

func (h *Handler) HandleCalc(w http.ResponseWriter, r *http.Request) {
	var req entity.CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, "Invalid request")
		return
	}
	resp := h.Endpoint.Calculate(context.Background(), req)
	if resp.Error != "" {
		respondWithError(w, resp.Error)
		return
	}
	respondWithResult(w, resp.Result)
}

func (h *Handler) HandleEvaluate(w http.ResponseWriter, r *http.Request) {
	var req entity.CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, "Invalid request")
		return
	}
	resp := h.Endpoint.Evaluate(context.Background(), req.Expr)
	if resp.Error != "" {
		respondWithError(w, resp.Error)
		return
	}
	respondWithResult(w, resp.Result)
}

func (h *Handler) HandleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func respondWithError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(entity.CalcResponse{Error: msg})
}

func respondWithResult(w http.ResponseWriter, res float64) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.CalcResponse{Result: res})
} 