package main

import (
	"bytes"
	"fmt"
	"log"

	"os"

	"github.com/el-pol/lovebot/fetch"
	"github.com/joho/godotenv"

	"github.com/dghubble/oauth1"
)

func main() {
	godotenv.Load()
	result, err := fetch.Fetch()

	if err != nil {
		log.Fatalf("Error when fetching: %v", err)
	}

	if result == "" {
		log.Fatalln("Result is empty")
	}

	if len(result) >= 280 {
		log.Fatalln("Result is too long")
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")
	accessToken := os.Getenv("ACCESS_TOKEN")
	accessSecret := os.Getenv("TOKEN_SECRET")

	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	path := "https://api.twitter.com/2/tweets"

	jsonBody := fmt.Sprintf(`{"text": "%s"}`, result)

	bodyReader := bytes.NewReader([]byte(jsonBody))
	resp, err := httpClient.Post(path, "application/json", bodyReader)
	if err != nil {
		log.Fatalf("Error when posting to twitter: %v", err)
	}
	defer resp.Body.Close()
	log.Printf("Response was OK: %v", resp)
}
