package app

import "github.com/flipped-aurora/gin-vue-admin/server/global"

/*
final String name;
final int quantity;
final Category category;
*/
type Grocery struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name"`
	Quantity string `json:"quantity" form:"quantity"`
	Category string `json:"category" form:"category"`
}

func (Grocery) TableName() string {
	return "groceries"
}
