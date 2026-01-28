// 自动生成模板Book
package app

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 书 结构体  Book
type Book struct {
	global.GVA_MODEL
	Book *string `json:"book" form:"book" gorm:"column:book;"` //书名
}

// TableName 书 Book自定义表名 books
func (Book) TableName() string {
	return "books"
}
