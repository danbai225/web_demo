package html

import (
	"github.com/gin-gonic/gin"
	"web_demo/internal/pkg/core"
)

// Index
// @Summary 返回index html
// @Description  返回index html
// @Tags API.html
// @Accept plain
// @Produce text/html
// @Success 200 string html
// @Failure 400 {object} code.Failure
// @Router /index [get]
func (h *handler) Index() core.HandlerFunc {
	return func(ctx core.Context) {
		ctx.HTML("index", gin.H{"msg": "hello"})
	}
}
