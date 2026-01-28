package app

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Demo struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name"`
}

func (Demo) TableName() string {
	return "demos"
}
