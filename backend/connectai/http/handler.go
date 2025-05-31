package http

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"context"
	"intelligent-calculator/connectai/entity"
	"intelligent-calculator/connectai/endpoint"
)

type Handler struct {
	Endpoint *endpoint.ConnectAIEndpoint
}

func NewHandler(ep *endpoint.ConnectAIEndpoint) *Handler {
	return &Handler{Endpoint: ep}
}

func (h *Handler) HandleAsk(w http.ResponseWriter, r *http.Request) {
	var req entity.ConnectAIRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, "Invalid request")
		return
	}
	promptType := r.Header.Get("X-Prompt-Type")
	if promptType == "" {
		promptType = "default"
	}
	resp := h.Endpoint.Ask(context.Background(), req, promptType)
	if resp.Error != "" {
		respondWithError(w, resp.Error)
		return
	}
	respondWithAnswer(w, resp.Answer)
}

func (h *Handler) HandleAskWithImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10MB max
	if err != nil {
		respondWithError(w, "Invalid multipart form")
		return
	}
	prompt := r.FormValue("prompt")
	promptType := r.Header.Get("X-Prompt-Type")
	if promptType == "" {
		promptType = "default"
	}
	file, _, err := r.FormFile("image")
	var imageBytes []byte
	if err == nil && file != nil {
		defer file.Close()
		imageBytes, _ = ioutil.ReadAll(file)
	}
	endpointReq := entity.ConnectAIImageRequest{Prompt: prompt, Image: imageBytes}
	resp := h.Endpoint.AskWithImage(context.Background(), endpointReq, promptType)
	if resp.Error != "" {
		respondWithError(w, resp.Error)
		return
	}
	respondWithAnswer(w, resp.Answer)
}

func respondWithError(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(entity.ConnectAIResponse{Error: msg})
}

func respondWithAnswer(w http.ResponseWriter, answer string) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(entity.ConnectAIResponse{Answer: answer})
} 