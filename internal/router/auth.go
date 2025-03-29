package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/auth"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func AuthRoutes(r *server.Hertz) {

	apiGroup := r.Group("/api")
	{
		apiGroup.Use(middleware.CheckAccessToken())
	}

	r.POST("/api/auth/emailVerify", auth.EmailVerify)
}
