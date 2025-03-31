package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/user"
	"github.com/cloudwego/hertz/pkg/route"
)

func NotJwtUserRoutes(r *route.RouterGroup) {
	r.POST("/user/sign", user.Sign)
	r.POST("/user/login", user.Login)
}

func JwtUserRoutes(r *route.RouterGroup) {
	r.GET("/api/user/userinfo", user.GetUser)
	r.PUT("/api/user/update", user.UpdateUser)
}
