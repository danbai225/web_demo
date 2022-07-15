package router

import (
	"web_demo/internal/api/admin"
	"web_demo/internal/api/html"
	"web_demo/internal/pkg/core"
)

func setApiRouter(r *resource) {
	//无须验证
	{
		group := r.mux.Group("")
		htmlHandler := html.New(r.logger, r.db, r.cache)
		group.GET("/index", htmlHandler.Index())
	}
	// 需要登录验证权限验证
	api := r.mux.Group("/api", r.interceptors.BaseInfo(), core.WrapAuthHandler(r.interceptors.CheckLogin))
	{
		{
			//admin
			adminGroup := api.Group("/admin", r.interceptors.Admin())
			adminHandler := admin.New(r.logger, r.db, r.cache)
			adminGroup.GET("/setting", adminHandler.SettingGet())
			adminGroup.POST("/setting", adminHandler.SettingPost())

		}
	}
}
