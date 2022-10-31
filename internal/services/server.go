package services

import (
	logs "github.com/danbai225/go-logs"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/services/base"
	"web_demo/internal/services/event"
	"web_demo/internal/services/user"
)

var SeverM *base.ServerManage

// InitService 注册服务
func InitService(db mysql.Repo, rdb redis.Repo) {
	SeverM = base.NewServerManage()
	defer func() {
		err := SeverM.Check()
		if err != nil {
			logs.Err(err)
		}
	}()
	SeverM.Set("user", user.New(SeverM, db, rdb))
	SeverM.Set("event", event.New(SeverM, db, rdb))
}
func Close() error {
	return SeverM.Over()
}
