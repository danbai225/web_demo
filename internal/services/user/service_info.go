package user

import (
	"web_demo/internal/core"
	"web_demo/internal/data_obj"
)

func (s *service) Info() (res *data_obj.UserInfoRes, err error) {
	ctx := core.GetContext()
	info := ctx.SessionUserInfo()
	s.eventLogServer().New("Info", "请求Info", nil)
	return &data_obj.UserInfoRes{
		Username: info.Username,
		Email:    info.Email,
		IP:       info.HttpBaseInfo.Ip,
	}, err
}
