package main

import (
	"fmt"
	"net/http"
	calculatorservice "intelligent-calculator/calculator/service"
	calculatorendpoint "intelligent-calculator/calculator/endpoint"
	calculatorhttp "intelligent-calculator/calculator/http"
	connectaiservice "intelligent-calculator/connectai/service"
	connectaiendpoint "intelligent-calculator/connectai/endpoint"
	connectaihttp "intelligent-calculator/connectai/http"
)

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		h(w, r)
	}
}

func main() {
	// Calculator domain
	calcService := calculatorservice.NewCalculatorService()
	calcEndpoint := calculatorendpoint.NewCalcEndpoint(calcService)
	calcHandler := calculatorhttp.NewHandler(calcEndpoint)

	// ConnectAI domain
	aiService := connectaiservice.NewConnectAIService()
	aiEndpoint := connectaiendpoint.NewConnectAIEndpoint(aiService)
	aiHandler := connectaihttp.NewHandler(aiEndpoint)

	http.HandleFunc("/api/calc", withCORS(calcHandler.HandleCalc))
	http.HandleFunc("/api/evaluate", withCORS(calcHandler.HandleEvaluate))
	http.HandleFunc("/api/connect-ai/ask", withCORS(aiHandler.HandleAsk))
	http.HandleFunc("/api/connect-ai/process-image", withCORS(aiHandler.HandleAskWithImage))
	http.HandleFunc("/health", calcHandler.HandleHealth)

	fmt.Println("Calculator backend running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
