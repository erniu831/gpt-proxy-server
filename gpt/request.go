package gpt

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"quick-talk/conf"
	"quick-talk/service/chat"
	"quick-talk/types"
	"quick-talk/util"
	"strings"
	"time"

	"errors"

	"github.com/gin-gonic/gin"
	gogpt "github.com/sashabaranov/go-openai"
	"golang.org/x/net/proxy"
)

var chatModels = []string{gogpt.GPT432K0314, gogpt.GPT4, gogpt.GPT40314, gogpt.GPT432K, gogpt.GPT3Dot5Turbo, gogpt.GPT3Dot5Turbo0301}

type ChatController struct {
}

func (c *ChatController) ResponseJson(ctx *gin.Context, code int, e string, r gin.H) {
	fmt.Println("res:", r)
	fmt.Println("err:", e)
}

// Completion 回复
func Completion(ctx *gin.Context, req chat.CompletionService) (gin.H, error) {
	request := gogpt.ChatCompletionRequest{
		Messages: req.Messages,
	}

	cnf := conf.Conf
	apiKey, _ := util.DecodeBase64ToString(cnf.ApiKey)
	fmt.Println("apikey:", apiKey)
	gptConfig := gogpt.DefaultConfig(apiKey)

	// if cnf.Proxy != "" {
	// 	transport := &http.Transport{}

	// 	if strings.HasPrefix(cnf.Proxy, "socks5h://") {
	// 		// 创建一个 DialContext 对象，并设置代理服务器
	// 		dialContext, err := newDialContext(cnf.Proxy[10:])
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		transport.DialContext = dialContext
	// 	} else {
	// 		// 创建一个 HTTP Transport 对象，并设置代理服务器
	// 		proxyUrl, err := url.Parse(cnf.Proxy)
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		transport.Proxy = http.ProxyURL(proxyUrl)
	// 	}
	// 	// 创建一个 HTTP 客户端，并将 Transport 对象设置为其 Transport 字段
	// 	gptConfig.HTTPClient = &http.Client{
	// 		Transport: transport,
	// 	}

	// }

	// 自定义gptConfig.BaseURL
	if cnf.ApiURL != "" {
		gptConfig.BaseURL = cnf.ApiURL
	}

	client := gogpt.NewClientWithConfig(gptConfig)
	if request.Messages[0].Role != "system" {
		newMessage := append([]gogpt.ChatCompletionMessage{
			{Role: "system", Content: cnf.BotDesc},
		}, request.Messages...)
		request.Messages = newMessage
	}

	// cnf.Model 是否在 chatModels 中
	if types.Contains(chatModels, cnf.Model) {
		request.Model = cnf.Model
		resp, err := client.CreateChatCompletion(ctx, request)
		if err != nil {
			return nil, err
		}
		return gin.H{
			"reply":    resp.Choices[0].Message.Content,
			"messages": append(request.Messages, resp.Choices[0].Message),
			"test":     resp,
		}, nil
	} else {
		prompt := ""
		for _, item := range request.Messages {
			prompt += item.Content + "/n"
		}
		prompt = strings.Trim(prompt, "/n")

		req := gogpt.CompletionRequest{
			Model:            cnf.Model,
			MaxTokens:        cnf.MaxTokens,
			TopP:             cnf.TopP,
			FrequencyPenalty: cnf.FrequencyPenalty,
			PresencePenalty:  cnf.PresencePenalty,
			Prompt:           prompt,
		}

		resp, err := client.CreateCompletion(ctx, req)
		if err != nil {
			return nil, err
		}

		return gin.H{
			"reply": resp.Choices[0].Text,
			"messages": append(request.Messages, gogpt.ChatCompletionMessage{
				Role:    "assistant",
				Content: resp.Choices[0].Text,
			}),
		}, nil
	}
}

// Completion 回复
func CompletionSSE(ctx *gin.Context, req chat.CompletionService) error {
	request := gogpt.ChatCompletionRequest{
		Messages: req.Messages,
	}

	cnf := conf.Conf
	apiKey, _ := util.DecodeBase64ToString(cnf.ApiKey)
	fmt.Println("apikey:", apiKey)
	gptConfig := gogpt.DefaultConfig(apiKey)

	// 自定义gptConfig.BaseURL
	if cnf.ApiURL != "" {
		gptConfig.BaseURL = cnf.ApiURL
	}

	client := gogpt.NewClientWithConfig(gptConfig)
	if request.Messages[0].Role != "system" {
		newMessage := append([]gogpt.ChatCompletionMessage{
			{Role: "system", Content: cnf.BotDesc},
		}, request.Messages...)
		request.Messages = newMessage
	}

	// cnf.Model 是否在 chatModels 中
	if !types.Contains(chatModels, cnf.Model) {
		cnf.Model = "gpt-3.5-turbo"
	}

	request.Model = cnf.Model
	stream, err := client.CreateChatCompletionStream(ctx, request)
	if err != nil {
		fmt.Println("createStreamErr:", stream)
		return err
	}
	for receivedResponse, streamErr := stream.Recv(); streamErr == nil; {
		if receivedResponse.ID != "" {

			fmt.Println("receivedResponse:", receivedResponse)
			byteData, _ := json.Marshal(receivedResponse)
			ctx.Writer.WriteString(string(byteData))
		}
		if streamErr != nil {
			fmt.Println("streamErr:", streamErr)
		}
	}

	_, streamErr := stream.Recv()
	if errors.Is(streamErr, io.EOF) {
		ctx.SSEvent("complete", "")

	}

	// 发送完成事件并结束SSE连接
	return nil
	// return gin.H{
	// 	"reply":    resp.Choices[0].Message.Content,
	// 	"messages": append(request.Messages, resp.Choices[0].Message),
	// 	"test":     resp,
	// }, nil

}

type dialContextFunc func(ctx context.Context, network, address string) (net.Conn, error)

func newDialContext(socks5 string) (dialContextFunc, error) {
	baseDialer := &net.Dialer{
		Timeout:   60 * time.Second,
		KeepAlive: 60 * time.Second,
	}

	if socks5 != "" {
		// split socks5 proxy string [username:password@]host:port
		var auth *proxy.Auth = nil

		if strings.Contains(socks5, "@") {
			proxyInfo := strings.SplitN(socks5, "@", 2)
			proxyUser := strings.Split(proxyInfo[0], ":")
			if len(proxyUser) == 2 {
				auth = &proxy.Auth{
					User:     proxyUser[0],
					Password: proxyUser[1],
				}
			}
			socks5 = proxyInfo[1]
		}

		dialSocksProxy, err := proxy.SOCKS5("tcp", socks5, auth, baseDialer)
		if err != nil {
			return nil, err
		}

		contextDialer, ok := dialSocksProxy.(proxy.ContextDialer)
		if !ok {
			return nil, err
		}

		return contextDialer.DialContext, nil
	} else {
		return baseDialer.DialContext, nil
	}
}
