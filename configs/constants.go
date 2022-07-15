package configs

import "time"

const (
	// MinGoVersion 最小 Go 版本
	MinGoVersion = 1.16

	// ProjectVersion 项目版本
	ProjectVersion = "v0.1"

	// ProjectName 项目名称
	ProjectName = "web_demo"

	// ProjectDomain 项目域名
	ProjectDomain = "127.0.0.1"

	// ProjectPort 项目端口
	ProjectPort = ":80"

	// ProjectAccessLogFile 项目访问日志存放文件
	ProjectAccessLogFile = "./logs/" + ProjectName + ".log"

	// HeaderLoginToken 登录验证 Token，Header 中传递的参数
	HeaderLoginToken = "SessionKey"

	// RedisKeyPrefixLoginUser Redis Key 前缀 - 登录用户信息
	RedisKeyPrefixLoginUser = ProjectName + ":login-user:"

	// ZhCN 简体中文 - 中国
	ZhCN = "zh-cn"

	// EnUS 英文 - 美国
	EnUS = "en-us"

	// MaxRequestsPerSecond 每秒最大请求量
	MaxRequestsPerSecond = 10000

	// LoginSessionTTL 登录有效期为 24 小时
	LoginSessionTTL = time.Hour * 24
)
