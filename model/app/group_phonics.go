package app

import "github.com/ditrue/go-template/global"

type GroupPhonic struct {
	global.GVA_MODEL
	Symbol    string `json:"symbol" form:"symbol" gorm:"column:symbol;"`
	AudioPath string `json:"audio_path" form:"audio_path" gorm:"column:audio_path;"`
}

func (GroupPhonic) TableName() string {
	return "group_phonics"
}
