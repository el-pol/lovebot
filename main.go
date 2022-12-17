package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"
)

func main() {
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
		Prompt:    "Generate a love advice sentence. It should be written in the style of an Indian guru:",
	}
	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(strings.TrimSpace(resp.Choices[0].Text))
}
