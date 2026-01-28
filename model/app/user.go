package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type Login interface {
	GetUsername() string
	GetUserId() uint
	GetUserInfo() any
}

type User struct {
	global.GVA_MODEL
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) GetUsername() string {
	return u.Username
}

func (u *User) GetUserId() uint {
	return u.ID
}

func (u *User) GetUserInfo() any {
	return u
}
