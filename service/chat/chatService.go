package chat

import gogpt "github.com/sashabaranov/go-openai"

type CompletionService struct {
	Messages []gogpt.ChatCompletionMessage `json:"messages"`
}

type CompletionMessage struct {
	Content string `json:"content" binding:"required"`
	Role    string `json:"role"`
}

type ChatProcessService struct {
	Prompt        string `json:"prompt"`
	Options       string `json:"options"`
	SystemMessage string `json:"systemMessage"`
}
