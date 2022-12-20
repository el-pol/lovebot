package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dghubble/oauth1"
	"github.com/el-pol/lovebot/fetch"
	"github.com/joho/godotenv"
)

func HandleRequest(ctx context.Context) (string, error) {
	godotenv.Load()

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("TOKEN_SECRET")
	prompt := os.Getenv("PROMPT")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	fetched := fetch.GetGenerated(prompt)

	// Post to Twitter
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/2/tweets"

	body := fmt.Sprintf(`{"text": "%s"}`, fetched)

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
