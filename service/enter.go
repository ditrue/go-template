package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/app"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	AppServiceGroup app.ServiceGroup
}
