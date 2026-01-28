package router

import (
	"github.com/ditrue/go-template/router/app"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	App app.RouterGroup
}
