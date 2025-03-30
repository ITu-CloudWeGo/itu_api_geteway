package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func GeneratedRegister(r *server.Hertz) {
	apiGroupWithCheckAccessToken := r.Group("/api")
	{
		apiGroupWithCheckAccessToken.Use(middleware.CheckAccessToken())
	}
	AuthRoutes(apiGroupWithCheckAccessToken)
	TagsRoutes(apiGroupWithCheckAccessToken)
}
