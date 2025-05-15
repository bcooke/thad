package llm

import (
	"fmt"
	"os"
)

type Provider string

const (
	ProviderOpenAI Provider = "openai"
	ProviderOllama Provider = "ollama"
)

// NewClient creates a new LLM client based on the provider and configuration
func NewClient(provider Provider, cfg ClientConfig) (LLMClient, error) {
	switch provider {
	case ProviderOpenAI:
		// If API key is not in config, try environment variable
		if cfg.APIKey == "" {
			cfg.APIKey = os.Getenv("OPENAI_API_KEY")
		}
		return NewOpenAIClient(cfg)

	case ProviderOllama:
		return NewOllamaClient(cfg), nil

	default:
		return nil, fmt.Errorf("unsupported LLM provider: %s", provider)
	}
}
