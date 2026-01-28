package app

import (
	"net/http"

	"github.com/ditrue/go-template/model/app"
	appResp "github.com/ditrue/go-template/model/app/response"
	"github.com/ditrue/go-template/model/common/response"
	"github.com/ditrue/go-template/utils"
	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (l *LoginApi) Login(c *gin.Context) {
	var loginReq struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := usersService.GetUserByUsername(loginReq.Username)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if user.Password != loginReq.Password {
		response.FailWithMessage("密码错误", c)
		return
	}

	l.TokenNext(c, user)
}

// TokenNext
func (l *LoginApi) TokenNext(c *gin.Context, user *app.User) {
	token, claims, err := utils.LoginToken(user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithData(appResp.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.ExpiresAt.Unix(),
	}, c)
}
