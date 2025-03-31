package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func GeneratedRegister(r *server.Hertz) {
	notApiGroup := r.Group("/api")
	apiGroupWithCheckAccessToken := r.Group("/api")
	{
		apiGroupWithCheckAccessToken.Use(middleware.CheckAccessToken())
	}
	AuthRoutes(apiGroupWithCheckAccessToken)
	LikesRoutes(apiGroupWithCheckAccessToken)
	JwtUserRoutes(apiGroupWithCheckAccessToken)

	NotJwtUserRoutes(notApiGroup)
}
