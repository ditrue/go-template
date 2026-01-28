package v1

import (
	"github.com/ditrue/go-template/api/v1/app"
	"github.com/ditrue/go-template/api/v1/web"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	WebApiGroup web.ApiGroup
	AppApiGroup app.ApiGroup
}
