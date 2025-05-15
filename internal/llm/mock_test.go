package llm

import (
	"testing"
)

func TestMockClient(t *testing.T) {
	tests := []struct {
		name           string
		responses      map[string]string
		defaultResp    string
		prompt         string
		expectedOutput string
	}{
		{
			name: "exact match response",
			responses: map[string]string{
				"test prompt": "test response",
			},
			defaultResp:    "default response",
			prompt:         "test prompt",
			expectedOutput: "test response",
		},
		{
			name: "default response",
			responses: map[string]string{
				"test prompt": "test response",
			},
			defaultResp:    "default response",
			prompt:         "unknown prompt",
			expectedOutput: "default response",
		},
		{
			name:           "empty responses",
			responses:      map[string]string{},
			defaultResp:    "default response",
			prompt:         "any prompt",
			expectedOutput: "default response",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewMockClient(tt.responses, tt.defaultResp)
			response, err := client.Complete(tt.prompt)
			if err != nil {
				t.Errorf("MockClient.Complete() error = %v", err)
				return
			}
			if response != tt.expectedOutput {
				t.Errorf("MockClient.Complete() = %v, want %v", response, tt.expectedOutput)
			}
		})
	}
}
