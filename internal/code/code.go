package code

import (
	_ "embed"

	"web_demo/configs"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	UrlSignError       = 10105
	AdminLoginError    = 20206
	NotAdministrator   = 31001
)

func Text(code int) string {
	lang := configs.Get().Language.Local

	if lang == configs.ZhCN {
		return zhCNText[code]
	}

	if lang == configs.EnUS {
		return enUSText[code]
	}

	return zhCNText[code]
}
