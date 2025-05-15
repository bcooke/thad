package llm

// MockClient is a mock implementation of LLMClient for testing
type MockClient struct {
	Responses       map[string]string
	DefaultResponse string
}

// NewMockClient creates a new mock client with optional predefined responses
func NewMockClient(responses map[string]string, defaultResponse string) *MockClient {
	return &MockClient{
		Responses:       responses,
		DefaultResponse: defaultResponse,
	}
}

// Complete implements the LLMClient interface
func (c *MockClient) Complete(prompt string) (string, error) {
	if response, ok := c.Responses[prompt]; ok {
		return response, nil
	}
	return c.DefaultResponse, nil
}
