package chat

import gogpt "github.com/sashabaranov/go-openai"

type CompletionService struct {
	Messages []gogpt.ChatCompletionMessage `json:"messages"`
}

type CompletionMessage struct {
	Content string `json:"content"`
	Role    string `json:"role"`
}
