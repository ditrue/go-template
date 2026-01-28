package response

import "github.com/ditrue/go-template/model/app"

type LoginResponse struct {
	User      *app.User `json:"user"`
	Token     string    `json:"token"`
	ExpiresAt int64     `json:"expiresAt"`
}
