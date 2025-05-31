package endpoint

import (
	"context"
	"intelligent-calculator/calculator/entity"
	"intelligent-calculator/calculator/service"
)

type CalcEndpoint struct {
	Service *service.CalculatorService
}

func NewCalcEndpoint(svc *service.CalculatorService) *CalcEndpoint {
	return &CalcEndpoint{Service: svc}
}

type CalcRequest = entity.CalcRequest

type CalcResponse = entity.CalcResponse

func (e *CalcEndpoint) Calculate(_ context.Context, req CalcRequest) CalcResponse {
	res, err := e.Service.Calculate(req)
	if err != nil {
		return CalcResponse{Error: err.Error()}
	}
	return CalcResponse{Result: res}
}

func (e *CalcEndpoint) Evaluate(_ context.Context, expr string) CalcResponse {
	res, err := e.Service.Evaluate(expr)
	if err != nil {
		return CalcResponse{Error: err.Error()}
	}
	return CalcResponse{Result: res}
} 