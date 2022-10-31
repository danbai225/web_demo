package user

import "web_demo/internal/core"

// UserInfo
// @Summary 用户基础信息
// @Description 用户基础信息
// @Tags API.user
// @Accept plain
// @Produce application/json
// @Success 200 {object} base.ReturnMsg{Data=data_obj.UserInfoRes} "返回数据"
// @Router /api/user/info [get]
func (h *handler) UserInfo() {
	ctx := core.GetContext()
	data, err := h.userServer.Info()
	if err != nil {
		ctx.AbortError(err)
	} else {
		ctx.Payload(data)
	}
}
