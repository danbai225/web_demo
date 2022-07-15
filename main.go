package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"web_demo/configs"
	"web_demo/internal/router"
	"web_demo/pkg/env"
	"web_demo/pkg/logger"
	"web_demo/pkg/shutdown"
	"web_demo/pkg/timeutil"

	"go.uber.org/zap"
)

func main() {
	// 初始化 access logger
	var accessLogger *zap.Logger
	var err error
	if env.Active().IsDev() {
		accessLogger, err = logger.NewJSONLogger(
			logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
			logger.WithTimeLayout(timeutil.CSTLayout),
			logger.WithFileRotationP(configs.ProjectAccessLogFile),
		)
	} else {
		accessLogger, err = logger.NewJSONLogger(
			logger.WithDisableConsole(),
			logger.WithField("domain", fmt.Sprintf("%s[%s]", configs.ProjectName, env.Active().Value())),
			logger.WithTimeLayout(timeutil.CSTLayout),
			logger.WithFileRotationP(configs.ProjectAccessLogFile),
		)
	}
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = accessLogger.Sync()
	}()

	// 初始化 HTTP 服务
	s, err := router.NewHTTPServer(accessLogger)
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Addr:    configs.ProjectPort,
		Handler: s.Mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			accessLogger.Fatal("http server startup err", zap.Error(err))
		}
	}()

	// 优雅关闭
	shutdown.NewHook().Close(
		// 关闭 http server
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				accessLogger.Error("server shutdown err", zap.Error(err))
			}
		},

		// 关闭 db
		func() {
			if s.Db != nil {
				if err := s.Db.DbClose(); err != nil {
					accessLogger.Error("dbw close err", zap.Error(err))
				}
			}
		},

		// 关闭 cache
		func() {
			if s.Cache != nil {
				if err := s.Cache.Close(); err != nil {
					accessLogger.Error("cache close err", zap.Error(err))
				}
			}
		},
	)
}
