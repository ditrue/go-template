package utils

import (
	"github.com/ditrue/go-template/model/app"
	"github.com/ditrue/go-template/model/app/request"
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
