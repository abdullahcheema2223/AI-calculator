package service

import (
	"fmt"
	"math"
	"strconv"
	"intelligent-calculator/calculator/entity"
)

type CalculatorService struct{}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (s *CalculatorService) Calculate(req entity.CalcRequest) (float64, error) {
	switch req.Operation {
	case "add":
		return req.A + req.B, nil
	case "subtract":
		return req.A - req.B, nil
	case "multiply":
		return req.A * req.B, nil
	case "divide":
		if req.B == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return req.A / req.B, nil
	case "sqrt":
		if req.A < 0 {
			return 0, fmt.Errorf("cannot take sqrt of negative number")
		}
		return math.Sqrt(req.A), nil
	case "power":
		return math.Pow(req.A, req.B), nil
	default:
		return 0, fmt.Errorf("unknown operation")
	}
}

// Very basic expression evaluator (no parentheses, left-to-right)
func (s *CalculatorService) Evaluate(expr string) (float64, error) {
	var res float64
	var lastOp byte = '+'
	var numStr string
	for i := 0; i < len(expr); i++ {
		c := expr[i]
		if c >= '0' && c <= '9' || c == '.' {
			numStr += string(c)
		} else if c == ' ' {
			continue
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, fmt.Errorf("invalid number: %s", numStr)
			}
			switch lastOp {
			case '+':
				res += num
			case '-':
				res -= num
			case '*':
				res *= num
			case '/':
				if num == 0 {
					return 0, fmt.Errorf("division by zero")
				}
				res /= num
			}
			lastOp = c
			numStr = ""
		}
	}
	if numStr != "" {
		num, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, fmt.Errorf("invalid number: %s", numStr)
		}
		switch lastOp {
		case '+':
			res += num
		case '-':
			res -= num
		case '*':
			res *= num
		case '/':
			if num == 0 {
				return 0, fmt.Errorf("division by zero")
			}
			res /= num
		}
	}
	return res, nil
} 