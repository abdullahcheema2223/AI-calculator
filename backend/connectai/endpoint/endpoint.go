package endpoint

import (
	"context"
	"intelligent-calculator/connectai/entity"
	"intelligent-calculator/connectai/service"
)

type ConnectAIEndpoint struct {
	Service *service.ConnectAIService
}

func NewConnectAIEndpoint(svc *service.ConnectAIService) *ConnectAIEndpoint {
	return &ConnectAIEndpoint{Service: svc}
}

type ConnectAIRequest = entity.ConnectAIRequest

type ConnectAIResponse = entity.ConnectAIResponse

type ConnectAIImageRequest = entity.ConnectAIImageRequest

func (e *ConnectAIEndpoint) Ask(_ context.Context, req ConnectAIRequest, promptType string) ConnectAIResponse {
	answer, err := e.Service.Ask(promptType, req.Prompt)
	if err != nil {
		return ConnectAIResponse{Error: err.Error()}
	}
	return ConnectAIResponse{Answer: answer}
}

func (e *ConnectAIEndpoint) AskWithImage(_ context.Context, req ConnectAIImageRequest, promptType string) ConnectAIResponse {
	answer, err := e.Service.AskWithImage(promptType, req.Prompt, req.Image)
	if err != nil {
		return ConnectAIResponse{Error: err.Error()}
	}
	return ConnectAIResponse{Answer: answer}
} 