package event_log

import (
	"time"

	"web_demo/internal/pkg/core"
	"web_demo/internal/repository/mysql/event_log"
)

func (s *service) New(ctx core.Context, Type, Content string) {
	info := ctx.SessionUserInfo()
	model := event_log.NewModel()
	model.Username = "test"
	model.Type = Type
	model.Content = Content
	model.CreatedAt = time.Now()
	model.Device = info.Device
	model.Ip = info.IP
	model.UserId = 1
	model.Location = info.Location
	_, err := model.Create(s.db.GetDb())
	if err != nil {
		s.logger.Error(err.Error())
	}
}
func (s *service) NewSystem(Type, Content string) {
	model := event_log.NewModel()
	model.Username = "系统日志"
	model.Type = Type
	model.Content = Content
	model.CreatedAt = time.Now()
	model.Device = "系统"
	_, err := model.Create(s.db.GetDb())
	if err != nil {
		s.logger.Error(err.Error())
	}
}
