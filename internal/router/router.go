package router

import (
	"web_demo/internal/pkg/core"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/router/interceptor"
	"web_demo/internal/services"
	"web_demo/pkg/env"
	"web_demo/pkg/errors"

	"go.uber.org/zap"
)

type resource struct {
	mux          core.Mux
	logger       *zap.Logger
	db           mysql.Repo
	cache        redis.Repo
	interceptors interceptor.Interceptor
}

type Server struct {
	Mux   core.Mux
	Db    mysql.Repo
	Cache redis.Repo
}

func NewHTTPServer(logger *zap.Logger) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger required")
	}

	r := new(resource)
	r.logger = logger
	// 初始化 DB
	dbRepo, err := mysql.New()
	if err != nil {
		logger.Fatal("new db err", zap.Error(err))
	}
	r.db = dbRepo

	// 初始化 Cache
	cacheRepo, err := redis.New()
	if err != nil {
		logger.Fatal("new cache err", zap.Error(err))
	}
	r.cache = cacheRepo
	var mux core.Mux
	if env.Active().IsDev() {
		mux, err = core.New(logger,
			core.WithEnableCors(),
			core.WithEnableRate(),
		)

	} else {
		mux, err = core.New(logger,
			core.WithEnableCors(),
			core.WithEnableRate(),
			core.WithDisableSwagger(),
		)
	}
	if err != nil {
		panic(err)
	}

	//初始化server
	services.InitService(logger, r.cache, r.db)
	r.mux = mux
	r.interceptors = interceptor.New(logger, r.cache, r.db)
	// 设置 API 路由
	setApiRouter(r)
	// 设置 Socket 路由
	//setSocketRouter(r)

	s := new(Server)
	s.Mux = mux
	s.Db = r.db
	s.Cache = r.cache

	return s, nil
}
