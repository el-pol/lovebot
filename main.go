package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/oauth1"
	"github.com/joho/godotenv"
)

func HandleRequest(ctx context.Context) (string, error) {
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")
	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("TOKEN_SECRET")
	prompt := os.Getenv("PROMPT")

	if apiKey == "" || consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	// Create the request body
	jsonBody := fmt.Sprintf(`{"prompt": "%s", "max_tokens": 256, "model": "text-davinci-003"}`, prompt)

	// Create the request
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer([]byte(jsonBody)))

	if err != nil {
		log.Fatalf("Error when creating request: %v", err)
	}

	// Add the headers
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s",
		apiKey))
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Error when sending request: %v", err)
	}

	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != 200 {
		log.Fatalf("Response was not OK: %v", resp)
	}

	// Parse the response
	var respBody struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	err = json.NewDecoder(resp.Body).Decode(&respBody)

	if err != nil {
		log.Fatalf("Error when decoding response: %v", err)
	}

	trimmed := strings.TrimSpace(respBody.Choices[0].Text)

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

	body := fmt.Sprintf(`{"text": "%s"}`, trimmed)

	bodyReader := strings.NewReader(body)

	response, err := httpClient.Post(path, "application/json", bodyReader)

	if err != nil {
		log.Fatalf("Error when posting to twitter: %v", err)
	}

	defer response.Body.Close()
	log.Printf("Response was OK: %v", response)
	return "finished", nil
}

func main() {
	lambda.Start(HandleRequest)
}
