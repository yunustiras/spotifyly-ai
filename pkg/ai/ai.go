package ai

import (
	"context"
	"fmt"
	"spotifyly-ai/internal/models"

	openai "github.com/sashabaranov/go-openai"
)

type AIClient struct {
	client *openai.Client
}

func NewAIClient() (*AIClient, error) {
	// Initialize OpenAI client
	client := openai.NewClient("your-openai-api-key")
	return &AIClient{client: client}, nil
}

func (a *AIClient) GroupSongs(songs []models.Song, criteria string) (map[string][]models.Song, error) {
	// Prepare the prompt for AI
	prompt := fmt.Sprintf("Group the following songs by %s:\n", criteria)
	for _, song := range songs {
		prompt += fmt.Sprintf("- %s by %s\n", song.Name, song.Artist)
	}

	// Use resp to parse the AI response
	resp, err := a.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get AI response: %v", err)
	}

	// Use resp to parse the AI response
	fmt.Println(resp.Choices[0].Message.Content)

	// Parse the response and group songs
	groupedSongs := make(map[string][]models.Song)
	// Implement the logic to parse the AI response and group songs
	// This is a mock implementation
	groupedSongs["Pop"] = []models.Song{
		{Name: "Shape of You", Artist: "Ed Sheeran"},
		{Name: "Blinding Lights", Artist: "The Weeknd"},
	}
	groupedSongs["Rock"] = []models.Song{
		{Name: "Bohemian Rhapsody", Artist: "Queen"},
	}

	return groupedSongs, nil
}
