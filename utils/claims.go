package utils

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app/request"
)

func LoginToken(user app.Login) (token string, claims request.CustomClaims, err error) {
	j := NewJWT()
	claims = j.CreateClaims(request.BaseClaims{
		ID:       user.GetUserId(),
		Username: user.GetUsername(),
	})
	token, err = j.CreateToken(claims)
	return
}
