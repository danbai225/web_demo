package services

import (
	"go.uber.org/zap"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"
	"web_demo/internal/services/admin"
	"web_demo/internal/services/base"
	"web_demo/internal/services/event_log"
)

var AdminServer admin.Service
var EventLogServer event_log.Service
var SeverM *base.ServerManage

func InitService(logger *zap.Logger, cache redis.Repo, db mysql.Repo) {
	SeverM = base.ServerManage{}.New()
	EventLogServer = event_log.New(logger, db, cache)
	AdminServer = admin.New(logger, db, cache, SeverM)

	SeverM.Set("event_log", EventLogServer)
	SeverM.Set("admin", AdminServer)

	EventLogServer.CheckSever()
	AdminServer.CheckSever()
}
