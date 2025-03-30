package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/auth"
	"github.com/cloudwego/hertz/pkg/route"
)

func AuthRoutes(r *route.RouterGroup) {

	r.POST("/auth/emailVerify", auth.EmailVerify)
}
