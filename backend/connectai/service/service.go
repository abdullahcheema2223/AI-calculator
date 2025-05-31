package service

import (
	"bytes"
	"encoding/json"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type ConnectAIService struct {
	Prompts map[string]string
	APIKey  string
	mu      sync.Mutex
}

func NewConnectAIService() *ConnectAIService {
	// Load .env file
	godotenv.Load(filepath.Join(".env"))
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		panic("GEMINI_API_KEY environment variable not set")
	}
	// Load prompts.yaml
	promptsPath := filepath.Join("prompts.yaml")
	promptsBytes, err := ioutil.ReadFile(promptsPath)
	if err != nil {
		panic("Failed to read prompts.yaml: " + err.Error())
	}
	var prompts map[string]string
	if err := yaml.Unmarshal(promptsBytes, &prompts); err != nil {
		panic("Failed to parse prompts.yaml: " + err.Error())
	}
	return &ConnectAIService{
		Prompts: prompts,
		APIKey:  apiKey,
	}
}

func (s *ConnectAIService) Ask(promptType, prompt string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	systemPrompt, ok := s.Prompts[promptType]
	if !ok {
		systemPrompt = s.Prompts["default"]
	}
	fullPrompt := systemPrompt + "\n" + prompt

	url := "https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent?key=" + s.APIKey
	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{{"text": fullPrompt}}},
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Gemini API error: %s", string(respBytes))
	}

	var geminiResp struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.Unmarshal(respBytes, &geminiResp); err != nil {
		return "", err
	}
	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("No answer from Gemini")
	}
	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}

func (s *ConnectAIService) AskWithImage(promptType, prompt string, image []byte) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	systemPrompt, ok := s.Prompts[promptType]
	if !ok {
		systemPrompt = s.Prompts["default"]
	}
	priorityInstruction := "INSTRUCTION: The following user prompt has the highest priority. Always do what the user prompt says, even if the image suggests otherwise."
	fullPrompt := systemPrompt + "\n" + priorityInstruction + "\n" + prompt

	url := "https://generativelanguage.googleapis.com/v1/models/gemini-1.5-flash:generateContent?key=" + s.APIKey

	// Gemini expects base64-encoded image in the request
	imgBase64 := ""
	if len(image) > 0 {
		imgBase64 = "data:image/jpeg;base64," + encodeToBase64(image)
	}

	parts := []map[string]interface{}{
		{"text": fullPrompt},
	}
	if imgBase64 != "" {
		parts = append(parts, map[string]interface{}{"inline_data": map[string]string{"mime_type": "image/jpeg", "data": imgBase64[23:]}})
	}

	body := map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": parts},
		},
	}
	jsonBody, _ := json.Marshal(body)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Gemini API error: %s", string(respBytes))
	}

	var geminiResp struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}
	if err := json.Unmarshal(respBytes, &geminiResp); err != nil {
		return "", err
	}
	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("No answer from Gemini")
	}
	return geminiResp.Candidates[0].Content.Parts[0].Text, nil
}

func encodeToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
} 