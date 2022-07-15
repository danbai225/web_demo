package interceptor

import (
	"net/http"

	"web_demo/internal/code"
	"web_demo/internal/pkg/core"
)

func (i *interceptor) Admin() core.HandlerFunc {
	return func(c core.Context) {
		info := c.SessionUserInfo()
		if info.UserType == 1 {
			return
		}
		c.AbortWithError(core.Error(
			http.StatusUnauthorized,
			code.AuthorizationError,
			code.Text(code.NotAdministrator)))
	}
}
