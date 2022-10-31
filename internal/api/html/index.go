package html

import (
	"github.com/gin-gonic/gin"
	"web_demo/internal/core"
)

// Index 首页
func (h *handler) Index() {
	ctx := core.GetContext()
	ctx.HTML("index", gin.H{"msg": "hahah"})
}
