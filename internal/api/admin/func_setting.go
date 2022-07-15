package admin

import (
	"net/http"

	"web_demo/internal/api/base"
	"web_demo/internal/code"
	"web_demo/internal/pkg/core"
	"web_demo/internal/pkg/validation"
)

type SettingGetRequest struct {
	Key string `json:"key" form:"key" binding:"required"`
}

// SettingGet
// @Summary 获取设置
// @Description 获取设置
// @Tags API.admin
// @Accept plain
// @Produce application/json
// @Param key query string true "键"
// @Success 200 {object} base.ReturnMsg "返回数据"
// @Failure 400 {object} code.Failure
// @Router /api/admin/setting [get]
func (h *handler) SettingGet() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(SettingGetRequest)
		res := new(base.ReturnMsg)
		if err := ctx.ShouldBindQuery(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
		data, err := h.adminServer.SettingGet(ctx, req.Key)
		if err != nil {
			res.Code = 1
			res.Msg = err.Error()
		} else {
			res.Data = data
		}
		ctx.Payload(res)
	}
}

type SettingPostRequest struct {
	Key string `json:"key"  binding:"required"` //键名
	Val string `json:"val" binding:"required"`  //键值
}

// SettingPost
// @Summary 添加设置
// @Description 添加设置 一些系统设置保存接口
// @Tags API.admin
// @Accept application/json
// @Produce json
// @Param Request body SettingPostRequest true "请求信息"
// @Success 200 {object} base.ReturnMsg
// @Failure 400 {object} code.Failure
// @Router /api/admin/setting [post]
func (h *handler) SettingPost() core.HandlerFunc {
	return func(ctx core.Context) {
		req := new(SettingPostRequest)
		res := new(base.ReturnMsg)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ctx.AbortWithError(core.Error(
				http.StatusBadRequest,
				code.ParamBindError,
				validation.Error(err)).WithError(err),
			)
			return
		}
		err := h.adminServer.SettingPost(ctx, req.Key, req.Val)
		if err != nil {
			res.Code = 1
			res.Msg = err.Error()
		}
		ctx.Payload(res)
	}
}
