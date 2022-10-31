package interceptor

import (
	"fmt"
	"github.com/pkg/errors"
	"web_demo/config"
	"web_demo/internal/core"
	"web_demo/internal/myconst"
	"web_demo/pkg/env"
)

func CheckLogin() {
	ctx := core.GetContext()
	token, ok := ctx.GetGin().GetQuery("token")
	if !ok {
		token = ctx.GetHeader("token")
	}
	username, _ := ctx.GetGin().GetQuery("u")
	if env.Active().IsDev() && username == "" && token == "" {
		username = config.Get().Dev.Username
	}
	HttpBaseInfo := ctx.SessionUserInfo().HttpBaseInfo
	if token == "" && username == "" {
		//无效
		ctx.AbortError(errors.New("签名验证未通过"))
		return
	} else if username != "" {
		//dev
		ctx.SetSessionUserInfo(&core.SessionUserInfo{
			Username:     username,
			Email:        "test@qq.com",
			HttpBaseInfo: HttpBaseInfo,
		})
	} else {
		//正常
		info := core.SessionUserInfo{}
		err := ctx.GetRdb().GetOBJ(fmt.Sprint(myconst.TokenPrefix, token), &info)
		if err != nil {
			ctx.AbortError(err)
			return
		}
		info.HttpBaseInfo = HttpBaseInfo
		ctx.SetSessionUserInfo(&info)
	}
}
