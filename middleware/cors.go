package middleware

import (
	"regexp"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	// if gin.Mode() == gin.ReleaseMode {
	// 	// 生产环境需要配置跨域域名，否则403
	// config.AllowOrigins = []string{"http://192.168.0.100:1002/", "http://43.159.46.132"}
	// } else {
	// 	// 测试环境下模糊匹配本地开头的请求
	config.AllowOriginFunc = func(origin string) bool {
		if regexp.MustCompile(`^https?://43\.159\.46\.132(:\d+)?$`).MatchString(origin) {
			return true
		}

		if regexp.MustCompile(`^https?://(www\.|api\.)?cheap-ai.(com|xyz)(:\d+)?$`).MatchString(origin) {
			return true
		}
		if regexp.MustCompile(`^http://127\.0\.0\.1:\d+$`).MatchString(origin) {
			return true
		}
		if regexp.MustCompile(`http://192.168.0.100:1002`).MatchString(origin) {
			return true
		}
		if regexp.MustCompile(`^http://localhost:\d+$`).MatchString(origin) {
			return true
		}
		return false
	}

	// config.AllowAllOrigins = true
	config.AllowCredentials = true
	return cors.New(config)
}
