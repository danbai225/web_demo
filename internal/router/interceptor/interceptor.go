package interceptor

import (
	"github.com/oschwald/geoip2-golang"
	"web_demo/internal/pkg/core"
	"web_demo/internal/proposal"
	"web_demo/internal/repository/mysql"
	"web_demo/internal/repository/redis"

	"go.uber.org/zap"
)

var _ Interceptor = (*interceptor)(nil)

type Interceptor interface {
	// CheckLogin 验证是否登录
	CheckLogin(ctx core.Context) (info proposal.SessionUserInfo, err core.BusinessError)
	//Admin 管理员身份验证
	Admin() core.HandlerFunc
	//BaseInfo 获取设备信息ip地址信息
	BaseInfo() core.HandlerFunc
	// i 为了避免被其他包实现
	i()
}

type interceptor struct {
	logger *zap.Logger
	cache  redis.Repo
	db     mysql.Repo
	ipDB   *geoip2.Reader
}

func New(logger *zap.Logger, cache redis.Repo, db mysql.Repo) Interceptor {
	ipDB, err := geoip2.Open("assets/GeoLite2-City.mmdb")
	if err != nil {
		logger.Error(err.Error())
	}
	return &interceptor{
		logger: logger,
		cache:  cache,
		db:     db,
		ipDB:   ipDB,
	}
}

func (i *interceptor) i() {}
