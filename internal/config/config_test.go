package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.Model.Provider != "ollama" {
		t.Errorf("expected default provider to be 'ollama', got %q", cfg.Model.Provider)
	}
	if cfg.Model.Model != "codellama" {
		t.Errorf("expected default model to be 'codellama', got %q", cfg.Model.Model)
	}
	if cfg.Model.BaseURL != "http://localhost:11434" {
		t.Errorf("expected default base URL to be 'http://localhost:11434', got %q", cfg.Model.BaseURL)
	}
	if cfg.PromptPreamble == "" {
		t.Error("expected default prompt preamble to be non-empty")
	}
}

func TestLoadConfig(t *testing.T) {
	// Create a temporary directory for test config
	tmpDir, err := os.MkdirTemp("", "thad-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Set HOME to temp dir
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	defer os.Setenv("HOME", oldHome)

	// Create config directory
	configDir := filepath.Join(tmpDir, ".config", "thad")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		t.Fatalf("failed to create config dir: %v", err)
	}

	// Test loading with no config file (should return defaults)
	cfg, err := Load()
	if err != nil {
		t.Errorf("Load() with no config file returned error: %v", err)
	}
	if cfg.Model.Provider != "ollama" {
		t.Errorf("expected default provider, got %q", cfg.Model.Provider)
	}

	// Test loading with invalid config file
	invalidConfig := []byte("invalid: yaml: content")
	if err := os.WriteFile(filepath.Join(configDir, "config.yaml"), invalidConfig, 0644); err != nil {
		t.Fatalf("failed to write invalid config: %v", err)
	}
	if _, err := Load(); err == nil {
		t.Error("Load() with invalid config file should return error")
	}

	// Test loading with valid config file
	validConfig := []byte(`
model:
  provider: openai
  api_key: test-key
  model: gpt-4
prompt_preamble: "test preamble"
`)
	if err := os.WriteFile(filepath.Join(configDir, "config.yaml"), validConfig, 0644); err != nil {
		t.Fatalf("failed to write valid config: %v", err)
	}
	cfg, err = Load()
	if err != nil {
		t.Errorf("Load() with valid config file returned error: %v", err)
	}
	if cfg.Model.Provider != "openai" {
		t.Errorf("expected provider 'openai', got %q", cfg.Model.Provider)
	}
	if cfg.Model.APIKey != "test-key" {
		t.Errorf("expected API key 'test-key', got %q", cfg.Model.APIKey)
	}
	if cfg.Model.Model != "gpt-4" {
		t.Errorf("expected model 'gpt-4', got %q", cfg.Model.Model)
	}
	if cfg.PromptPreamble != "test preamble" {
		t.Errorf("expected prompt preamble 'test preamble', got %q", cfg.PromptPreamble)
	}
}
