package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/app"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	App app.RouterGroup
}
