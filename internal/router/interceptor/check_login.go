package interceptor

import (
	"encoding/json"
	"net/http"

	"web_demo/configs"
	"web_demo/internal/code"
	"web_demo/internal/pkg/core"
	"web_demo/internal/proposal"
	"web_demo/internal/repository/redis"
	"web_demo/pkg/env"
	"web_demo/pkg/errors"
)

func (i *interceptor) CheckLogin(ctx core.Context) (sessionUserInfo proposal.SessionUserInfo, err core.BusinessError) {
	val := ctx.GetVal("baseInfo")
	if val != nil {
		base := val.(*BaseInfo)
		defer func() {
			sessionUserInfo.Device = base.Device
			sessionUserInfo.IP = base.Ip
			sessionUserInfo.Location = base.Location
		}()
	}
	token := ctx.GetHeader(configs.HeaderLoginToken)
	if token == "" {
		token, _ = ctx.GetGin().GetQuery(configs.HeaderLoginToken)
	}
	if token == "" {
		if env.Active().IsDev() {
			sessionUserInfo = proposal.SessionUserInfo{
				UserType: 1,
				Device:   "test",
				IP:       "127.0.0.1",
				Location: "test",
			}
			return
		}
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("Header 中缺少 Session_Key 参数"))
		return
	}
	if !i.cache.Exists(configs.RedisKeyPrefixLoginUser + token) {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(errors.New("请先登录"))
		return
	}

	cacheData, cacheErr := i.cache.Get(configs.RedisKeyPrefixLoginUser+token, redis.WithTrace(ctx.Trace()))
	if cacheErr != nil {
		err = core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(cacheErr)

		return
	}
	i.cache.Expire(configs.RedisKeyPrefixLoginUser+token, configs.LoginSessionTTL)
	jsonErr := json.Unmarshal([]byte(cacheData), &sessionUserInfo)
	if jsonErr != nil {
		core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.AuthorizationError)).WithError(jsonErr)

		return
	}
	return
}
