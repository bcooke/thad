package llm

import (
	"testing"
)

func TestNewClient(t *testing.T) {
	tests := []struct {
		name     string
		provider Provider
		cfg      ClientConfig
		wantErr  bool
	}{
		{
			name:     "valid OpenAI config",
			provider: ProviderOpenAI,
			cfg: ClientConfig{
				APIKey:         "test-key",
				PromptPreamble: "test preamble",
				Model:          "gpt-4",
			},
			wantErr: false,
		},
		{
			name:     "OpenAI without API key",
			provider: ProviderOpenAI,
			cfg: ClientConfig{
				PromptPreamble: "test preamble",
			},
			wantErr: true,
		},
		{
			name:     "valid Ollama config",
			provider: ProviderOllama,
			cfg: ClientConfig{
				PromptPreamble: "test preamble",
				Model:          "codellama",
				BaseURL:        "http://localhost:11434",
			},
			wantErr: false,
		},
		{
			name:     "invalid provider",
			provider: "invalid",
			cfg: ClientConfig{
				PromptPreamble: "test preamble",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.provider, tt.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && client == nil {
				t.Error("NewClient() returned nil client when no error expected")
			}
		})
	}
}
