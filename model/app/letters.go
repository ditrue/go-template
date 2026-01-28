package app

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Letter struct {
	global.GVA_MODEL
	// 小写
	Lowercase string `json:"lowercase" form:"lowercase" gorm:"column:lowercase;"`
	// 大写
	Uppercase string `json:"uppercase" form:"uppercase" gorm:"column:uppercase;"`
	// 音频路基
	UppercaseAudioPath string `json:"uppercase_audio_path" form:"uppercase_audio_path" gorm:"column:uppercase_audio_path;"`
	LowercaseAudioPath string `json:"lowercase_audio_path" form:"lowercase_audio_path" gorm:"column:lowercase_audio_path;"`
}

func (Letter) TableName() string {
	return "letters"
}
