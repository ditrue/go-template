package app

import "github.com/ditrue/go-template/global"

type Demo struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name"`
}

func (Demo) TableName() string {
	return "demos"
}
