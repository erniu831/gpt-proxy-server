package server

import (
	"os"
	"quick-talk/api"
	"quick-talk/middleware"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		// 发券
		v1.POST("code/generate", api.CodeGenerate)

		// 查券
		v1.POST("code/check", api.CodeCheck)
		// 需要登录保护的
		auth := v1.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.POST("user/logout", api.UserLogout)
			// user
			auth.POST("session", api.GetSession)
			// 核券
			auth.POST("code/redeem", api.CodeRedeem)
		}

		// chat
		v1.POST("chat/completion", api.ChatCompletion)
	}

	sseV1 := r.Group("/sse/v1")
	sseV1.Use(middleware.AuthRequired())
	{
		sseV1.POST("chat/completion", api.ChatCompletionSSE)
	}

	sseOpen := r.Group("/sse/open")
	{
		sseOpen.POST("chat/completion", api.ChatCompletionForFreeSSE)
	}

	return r
}
