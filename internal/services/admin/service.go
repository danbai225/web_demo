package admin

import (
	"time"
	"web_demo/internal/pkg/core"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/services/base"

	"go.uber.org/zap"
)

var _ Service = (*service)(nil)

type eventServer interface {
	New(ctx core.Context, Type, Content string)
}
type Service interface {
	i()
	CheckSever()
	SettingGet(ctx core.Context, key string) (string, error)
	SettingPost(ctx core.Context, key, val string) error
}

type service struct {
	db                   mysql.Repo
	cache                redis.Repo
	logger               *zap.Logger
	EndOfThisPeriodTimer *time.Timer
	manage               *base.ServerManage
}

func New(logger *zap.Logger, db mysql.Repo, cache redis.Repo, manage *base.ServerManage) Service {
	s := &service{
		db:     db,
		cache:  cache,
		logger: logger,
		manage: manage,
	}
	return s
}

func (s *service) i() {}
func (s *service) CheckSever() {
	s.eventServer()
}
func (s *service) eventServer() eventServer {
	return s.manage.Get("event_log").(eventServer)
}
