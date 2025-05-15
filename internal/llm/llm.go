package llm

// LLMClient defines the interface for LLM providers
type LLMClient interface {
	// Complete takes a prompt and returns the model's response
	Complete(prompt string) (string, error)
}

// ClientConfig holds the configuration for an LLM client
type ClientConfig struct {
	PromptPreamble string
	Model          string
	BaseURL        string
	APIKey         string
}

// OpenAIClient and OllamaClient will implement LLMClient
