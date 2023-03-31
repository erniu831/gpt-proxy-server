package chat

import gogpt "github.com/sashabaranov/go-openai"

type CompletionService struct {
	Messages []gogpt.ChatCompletionMessage `json:"messages"`
}

type CompletionMessage struct {
	Content string `json:"content" binding:"required"`
	Role    string `json:"role"`
}
