package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type OllamaClient struct {
	baseURL        string
	model          string
	httpClient     *http.Client
	promptPreamble string
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
}

func NewOllamaClient(cfg ClientConfig) *OllamaClient {
	baseURL := cfg.BaseURL
	if baseURL == "" {
		baseURL = "http://localhost:11434"
	}

	model := cfg.Model
	if model == "" {
		model = "codellama" // Default to CodeLlama for command generation
	}

	return &OllamaClient{
		baseURL:        baseURL,
		model:          model,
		httpClient:     &http.Client{},
		promptPreamble: cfg.PromptPreamble,
	}
}

func (c *OllamaClient) Complete(prompt string) (string, error) {
	fullPrompt := fmt.Sprintf("%s\n\nUser: %s\nAssistant: Return only the shell command, nothing else:", c.promptPreamble, prompt)

	reqBody := ollamaRequest{
		Model:  c.model,
		Prompt: fullPrompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.httpClient.Post(
		fmt.Sprintf("%s/api/generate", c.baseURL),
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", fmt.Errorf("failed to make request to Ollama: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ollama API returned non-200 status code: %d", resp.StatusCode)
	}

	var ollamaResp ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", fmt.Errorf("failed to decode Ollama response: %w", err)
	}

	return ollamaResp.Response, nil
}
