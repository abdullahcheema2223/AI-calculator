package entity

type ConnectAIRequest struct {
	Prompt string `json:"prompt"`
}

type ConnectAIResponse struct {
	Answer string `json:"answer"`
	Error  string `json:"error,omitempty"`
}

type ConnectAIImageRequest struct {
	Prompt string
	Image  []byte
} 