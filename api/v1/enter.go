package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/app"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/web"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	WebApiGroup web.ApiGroup
	AppApiGroup app.ApiGroup
}
