package app

import (
	api "github.com/ditrue/go-template/api/v1"
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
