package app

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type Share struct {
	global.GVA_MODEL
	BookID   uint  `json:"-" form:"book_id"`
	Book     *Book `json:"book" form:"book"`
	ToUserID uint  `json:"to_user_id" form:"to_user_id"`
	UserID   uint  `json:"user_id" form:"user_id"`
}
