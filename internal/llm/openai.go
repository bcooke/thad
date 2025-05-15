package llm

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

type OpenAIClient struct {
	client         *openai.Client
	promptPreamble string
	model          string
}

func NewOpenAIClient(cfg ClientConfig) (*OpenAIClient, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("OpenAI API key is required")
	}

	client := openai.NewClient(cfg.APIKey)
	model := cfg.Model
	if model == "" {
		model = openai.GPT4TurboPreview
	}

	return &OpenAIClient{
		client:         client,
		promptPreamble: cfg.PromptPreamble,
		model:          model,
	}, nil
}

func (c *OpenAIClient) Complete(prompt string) (string, error) {
	ctx := context.Background()

	resp, err := c.client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: c.model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: c.promptPreamble,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
			Temperature: 0.3, // Lower temperature for more deterministic command generation
		},
	)

	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI API")
	}

	return resp.Choices[0].Message.Content, nil
}
