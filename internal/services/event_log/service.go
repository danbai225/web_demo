package event_log

import (
	"web_demo/internal/pkg/core"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"

	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type Service interface {
	i()
	CheckSever()
	New(ctx core.Context, Type, Content string)
	NewSystem(Type, Content string)
}

type service struct {
	db     mysql.Repo
	cache  redis.Repo
	logger *zap.Logger
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo) Service {
	return &service{
		db:     db,
		cache:  cache,
		logger: logger,
	}
}

func (s *service) i() {}
func (s *service) CheckSever() {

}
