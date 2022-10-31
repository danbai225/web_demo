package router

import (
	"web_demo/internal/api/html"
	"web_demo/internal/api/user"
	"web_demo/internal/core"
	"web_demo/internal/router/interceptor"
)

func RegRouter(g core.Mux) {
	//Handler
	userHandler := user.New()
	htmlHandler := html.New()
	rootGroup := g.Group("/", interceptor.Cors, interceptor.BaseInfo)
	//无需登陆
	{
		rootGroup.GET("/", htmlHandler.Index)
		rootGroup.GET("/index", htmlHandler.Index)
	}
	//登陆
	{
		apiGroup := rootGroup.Group("/api", interceptor.CheckLogin)
		userGroup := apiGroup.Group("/user")
		userGroup.GET("/info", userHandler.UserInfo)
	}
}
