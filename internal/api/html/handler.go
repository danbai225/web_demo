package html

import (
	"go.uber.org/zap"
	"web_demo/internal/pkg/core"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/services"
	"web_demo/internal/services/admin"
)

var _ Handler = (*handler)(nil)

type Handler interface {
	i()
	// Index index html
	// @Tags API.html
	// @Router /index [get]
	Index() core.HandlerFunc
}

type handler struct {
	logger      *zap.Logger
	db          mysql.Repo
	cache       redis.Repo
	adminServer admin.Service
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Handler {
	return &handler{
		logger:      logger,
		db:          db,
		cache:       cache,
		adminServer: services.AdminServer,
	}
}

func (h *handler) i() {}
