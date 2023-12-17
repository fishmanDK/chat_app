package transform

import (
	fishman_auth_v1 "github.com/fishmanDK/proto-auth/gen/go/auth"
	"realtime_chat_app"
)

func UserTransform(user realtime_chat_app.User) *fishman_auth_v1.User {
	return &fishman_auth_v1.User{
		AppName:  user.AppName,
		Email:    user.Email,
		Password: user.Password,
	}
}
