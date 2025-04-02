package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func GeneratedRegister(r *server.Hertz) {

	apiWithAccessJwt := r.Group("/api")
	{
		apiWithAccessJwt.Use(middleware.CheckAccessToken())
		AuthRoutes(apiWithAccessJwt)
		LikesRoutes(apiWithAccessJwt)
		JwtUserRoutes(apiWithAccessJwt)
		PostRoutes(apiWithAccessJwt)
	}

	notCheckedApi := r.Group("/api")
	{
		NotJwtUserRoutes(notCheckedApi)
	}

}
