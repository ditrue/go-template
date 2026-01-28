package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/app"

type LoginResponse struct {
	User      *app.User `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expiresAt"`
}
