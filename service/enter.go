package service

import (
	"github.com/ditrue/go-template/service/app"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	AppServiceGroup app.ServiceGroup
}
