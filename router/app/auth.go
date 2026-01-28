package app

import (
	api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct{}

var loginApi = api.ApiGroupApp.AppApiGroup.LoginApi

func (l *AuthRouter) InitAuthRouter(PublicRouter *gin.RouterGroup, PrivateRouter *gin.RouterGroup) {
	authRouter := PublicRouter.Group("auth")
	{
		authRouter.POST("login", loginApi.Login)
	}
}
