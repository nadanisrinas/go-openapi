package main

import (
	"bufio"
	"fmt"
	"openapi-golang/request"
	"os"
)

func main() {

	apiKey := ""
	apiUrl := "https://api.openai.com/v1/chat/completions"
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your message: ")
	scanner.Scan()
	userInput := scanner.Text()

	payload := fmt.Sprintf(`{
		"model": "gpt-3.5-turbo",
		"messages": [
			{
				"role": "user",
				"content": "%s"
			}
		] 
	}`, userInput)
	request.RequestPostA(apiUrl, payload, apiKey)

}
