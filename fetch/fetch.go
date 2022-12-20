package fetch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetGenerated(prompt string) string {
	godotenv.Load()

	apiKey := os.Getenv("OPENAI_API_KEY")

	if apiKey == "" {
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

	return trimmed
}
