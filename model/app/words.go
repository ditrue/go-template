package app

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Word struct {
	global.GVA_MODEL
	Word      string `json:"word" form:"word" gorm:"column:word;"`
	AudioPath string `json:"audio_path" form:"audio_path" gorm:"column:audio_path;"`
}

func (Word) TableName() string {
	return "words"
}

