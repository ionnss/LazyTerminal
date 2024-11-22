package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func saveTokenToEnv(token string) error {
	// Write the token to the .env file
	file, err := os.Create(".env")
	if err != nil {
		return fmt.Errorf("error creating .env file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("OPENAI_API_KEY=%s\n", token))
	if err != nil {
		return fmt.Errorf("error writing to .env file: %v", err)
	}

	return nil
}

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No .env file found. A new one will be created.")
	}

	// Get the API token from the .env file
	apiToken := os.Getenv("OPENAI_API_KEY")
	if apiToken == "" {
		// If the token does not exist, request it from the user
		fmt.Print("Enter your OpenAI API token: ")
		reader := bufio.NewReader(os.Stdin)
		apiToken, _ = reader.ReadString('\n')
		apiToken = strings.TrimSpace(apiToken)

		// Save the token to the .env file
		err := saveTokenToEnv(apiToken)
		if err != nil {
			fmt.Printf("Error saving token to .env file: %v\n", err)
			return
		}
		fmt.Println("Token saved to .env file.")
	}

	// Loop for action commands
	for {
		fmt.Print("Type a command (ex: lzt 'your question', or 'lzt -ut' to update token) or 'exit' to exit: ")
		reader := bufio.NewReader(os.Stdin)
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		// Check if command == exit
		if command == "exit" {
			fmt.Println("LaZzZzZzZy out...")
			break
		}

		// Check for token update command
		if command == "lzt -ut" {
			fmt.Print("Enter your new OpenAI API token: ")
			newToken, _ := reader.ReadString('\n')
			newToken = strings.TrimSpace(newToken)

			// Save the new token
			err := saveTokenToEnv(newToken)
			if err != nil {
				fmt.Printf("Error updating token in .env file: %v\n", err)
			} else {
				apiToken = newToken
				fmt.Println("Token successfully updated and saved to .env file.")
			}
			continue
		}

		// Check lzt command
		if strings.HasPrefix(command, "lzt ") {
			input := strings.TrimPrefix(command, "lzt ")
			response, err := GetGPTResponse(apiToken, input)
			if err != nil {
				fmt.Printf("Trouble communicating with API: %v\n", err)
			} else {
				fmt.Println("LazyTerminal >", response)
			}
		} else {
			fmt.Println("Bad command > Use lzt 'your question', or 'lzt -ut' to update token")
		}
	}
}

