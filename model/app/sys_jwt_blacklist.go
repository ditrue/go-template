package app

import (
	"github.com/ditrue/go-template/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
