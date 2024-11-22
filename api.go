package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const apiURL = "https://api.openai.com/v1/chat/completions"

// Structure for the API request body
type GPTRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
}

// Structure for the API response
type GPTResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error struct {
		Message string `json:"message"`
	} `json:"error"`
}

// Function to send the message and receive the GPT response
func GetGPTResponse(apiToken, input string) (string, error) {
	// Build the request body
	requestBody := GPTRequest{
		Model: "gpt-3.5-turbo",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "user", Content: input},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error creating JSON: %v", err)
	}

	// Configure the HTTP request
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating the request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiToken)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending the request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading the response: %v", err)
	}

	// Check the response status
	if resp.StatusCode != http.StatusOK {
		var apiError GPTResponse
		json.Unmarshal(body, &apiError)
		if apiError.Error.Message != "" {
			return "", errors.New("API error: " + apiError.Error.Message)
		}
		return "", errors.New("unknown API error")
	}

	// Parse the response
	var gptResp GPTResponse
	if err := json.Unmarshal(body, &gptResp); err != nil {
		return "", fmt.Errorf("error parsing the response: %v", err)
	}

	if len(gptResp.Choices) > 0 {
		return gptResp.Choices[0].Message.Content, nil
	}

	return "", errors.New("empty response from API")
}

