package chatgpt

import (
	"context"
	"log"

	gogpt "github.com/sashabaranov/go-openai"
)

func QueryGPT(APIKey string, QueryText string) string {

	log.Println("Querying ChatGPT....")

	c := gogpt.NewClient(APIKey)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:       gogpt.GPT3TextDavinci003,
		MaxTokens:   3000,
		Prompt:      QueryText,
		Temperature: 0,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		log.Println("Error while quering ChatGPT, ", err)
	}

	log.Println("ChatGPT querying complete")

	return resp.Choices[0].Text
}
