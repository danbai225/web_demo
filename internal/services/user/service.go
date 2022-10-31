package user

import (
	"web_demo/internal/data_obj"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/services/base"
)

var _ Service = (*service)(nil)

type Service interface {
	base.RegServer
	i()
	Info() (res *data_obj.UserInfoRes, err error)
}

type service struct {
	serverManage *base.ServerManage
	db           mysql.Repo
	rdb          redis.Repo
}

func New(serverManage *base.ServerManage, db mysql.Repo, rdb redis.Repo) Service {
	return &service{
		serverManage: serverManage,
		db:           db,
		rdb:          rdb,
	}
}

func (s *service) i() {}
func (s *service) CheckSever() error {
	s.eventLogServer()
	return nil
}
func (s *service) OverServer() error {
	return nil
}

type eventLogServer interface {
	New(Type, Content string, data interface{})
}

func (s *service) eventLogServer() eventLogServer {
	return s.serverManage.Get("event").(eventLogServer)
}
