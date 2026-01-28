package app

import (
	"fmt"

	"github.com/ditrue/go-template/global"
	"github.com/ditrue/go-template/model/app"
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
