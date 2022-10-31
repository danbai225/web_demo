package interceptor

import (
	"net/http"
	"web_demo/internal/core"
)

func Cors() {
	context := core.GetContext()
	ginContext := context.GetGin()
	method := ginContext.Request.Method
	ginContext.Header("Access-Control-Allow-Origin", "*")
	ginContext.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, Authorization, Token")
	ginContext.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	ginContext.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	ginContext.Header("Access-Control-Allow-Credentials", "true")
	// 允许放行OPTIONS请求
	if method == "OPTIONS" {
		ginContext.AbortWithStatus(http.StatusNoContent)
	}
	ginContext.Next()
}
