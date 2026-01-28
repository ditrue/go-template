package request

import "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"

type GrocerySearch struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	Keyword  string `json:"keyword" form:"keyword"`
}

type GroceryCreate struct {
	Name     string `json:"name" form:"name"`
	Quantity string `json:"quantity" form:"quantity"`
	Category string `json:"category" form:"category"`
}

type GetGroceries struct {
	// PageInfo
	request.PageInfo
}
