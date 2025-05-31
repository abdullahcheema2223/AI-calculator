package entity

type CalcRequest struct {
	A         float64 `json:"a,omitempty"`
	B         float64 `json:"b,omitempty"`
	Operation string  `json:"operation,omitempty"`
	Expr      string  `json:"expr,omitempty"`
}

type CalcResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
} 