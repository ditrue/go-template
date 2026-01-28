package app

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/app"
)

type UsersService struct{}

func (u *UsersService) GetUserByUsername(username string) (*app.User, error) {
	fmt.Println(username)
	var user app.User
	err := global.GVA_DB.First(&user, "username = ?", username).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
