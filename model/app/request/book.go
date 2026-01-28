package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BookSearch struct {
	CreatedAtRange []time.Time `json:"createdAtRange" form:"createdAtRange[]"`
	request.PageInfo
}

type CreateBook struct {
	ID       uint `json:"id" form:"id" uri:"id"`
	ToUserID uint `json:"to_user_id" form:"to_user_id"`
}
