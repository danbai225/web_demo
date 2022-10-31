package event

import (
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/services/base"
)

var _ Service = (*service)(nil)

type Service interface {
	base.RegServer
	i()
	New(Type, Content string, data interface{})
}

type service struct {
	serverManage *base.ServerManage
	db           mysql.Repo
}

func New(serverManage *base.ServerManage, db mysql.Repo, rdb redis.Repo) Service {
	return &service{
		serverManage: serverManage,
		db:           db,
	}
}

func (s *service) i() {}
func (s *service) CheckSever() error {
	return nil
}
func (s *service) OverServer() error {
	return nil
}
