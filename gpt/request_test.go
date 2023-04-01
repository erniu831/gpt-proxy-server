package gpt

import (
	"context"
	"fmt"
	"testing"

	gogpt "github.com/sashabaranov/go-openai"
)

func TestAdd(t *testing.T) {

	client := gogpt.NewClient("")
	res, err := client.CreateChatCompletion(context.Background(), gogpt.ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []gogpt.ChatCompletionMessage{
			{
				Content: "hello",
				Role:    "user",
			},
		},
	})
	fmt.Println("err:", err)
	fmt.Println("res:", res)
}
