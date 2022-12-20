package fetch

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func Fetch() (string, error) {
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
		log.Fatalln("Missing API KEY")
	}

	c := gogpt.NewClient(apiKey)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 256,
		Prompt:    os.Getenv("PROMPT"),
	}

	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		return "", err
	}

	log.Printf("Completion was successful: %+v", resp)

	trimmed := strings.TrimSpace(resp.Choices[0].Text)

	return trimmed, nil
}
