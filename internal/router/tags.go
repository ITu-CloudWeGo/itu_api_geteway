package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/tags"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func TagsRoutes(r *server.Hertz) {

	r.POST("/api/tag/tags", tags.CreateTags)
}
