package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/auth"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func AuthRoutes(r *server.Hertz) {
	r.POST("auth/refresh", auth.RefreshToken)
	r.POST("auth/emailVerify", auth.EmailVerify)
}
