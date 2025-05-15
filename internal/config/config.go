package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// DefaultConfig returns the default configuration
func DefaultConfig() *Config {
	return &Config{
		Model: ModelConfig{
			Provider:    "ollama", // Default to Ollama since it's free and local
			OllamaModel: "codellama",
			BaseURL:     "http://localhost:11434",
		},
		PromptPreamble: "You are an expert shell assistant. Return the shortest working command for the user's OS unless they ask for alternatives.",
	}
}

type ModelConfig struct {
	Provider    string `yaml:"provider"`
	APIKey      string `yaml:"api_key"`
	OpenAIModel string `yaml:"openai_model"`
	OllamaModel string `yaml:"ollama_model"`
	BaseURL     string `yaml:"base_url"`
}

type Config struct {
	Model          ModelConfig `yaml:"model"`
	PromptPreamble string      `yaml:"prompt_preamble"`
}

// Load loads the configuration from the config file, falling back to defaults if not found
func Load() (*Config, error) {
	cfg := DefaultConfig()

	cfgPath := filepath.Join(os.Getenv("HOME"), ".config", "thad", "config.yaml")
	f, err := os.Open(cfgPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Return default config if file doesn't exist
			return cfg, nil
		}
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer f.Close()

	dec := yaml.NewDecoder(f)
	if err := dec.Decode(cfg); err != nil {
		return nil, fmt.Errorf("failed to decode config file: %w", err)
	}

	// Validate required fields
	if cfg.Model.Provider == "" {
		return nil, fmt.Errorf("model provider is required")
	}

	return cfg, nil
}
