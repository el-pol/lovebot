package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"strings"

	"os"

	"github.com/joho/godotenv"
	gogpt "github.com/sashabaranov/go-gpt3"

	"github.com/dghubble/oauth1"
)

func main() {
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("TOKEN_SECRET")

	if apiKey == "" || consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	// Fetch from GPT-3
	c := gogpt.NewClient(apiKey)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 256,
		Prompt:    os.Getenv("PROMPT"),
	}

	resp, err := c.CreateCompletion(ctx, req)

	if err != nil {
		log.Fatalf("Error when creating completion: %v", err)
	}

	log.Printf("Completion was successful: %+v", resp)

	trimmed := strings.TrimSpace(resp.Choices[0].Text)

	if trimmed == "" {
		log.Fatalln("Result is empty")
	}

	if len(trimmed) >= 280 {
		log.Fatalln("Result is too long")
	}

	// Post to Twitter
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/2/tweets"

	jsonBody := fmt.Sprintf(`{"text": "%s"}`, trimmed)

	bodyReader := bytes.NewReader([]byte(jsonBody))

	response, err := httpClient.Post(path, "application/json", bodyReader)

	if err != nil {
		log.Fatalf("Error when posting to twitter: %v", err)
	}

	defer response.Body.Close()
	log.Printf("Response was OK: %v", response)
}
