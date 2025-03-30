package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/tgas"
	"github.com/cloudwego/hertz/pkg/route"
)

func TagsRoutes(r *route.RouterGroup) {
	r.POST("/tag/tags", tgas.CreateTags)
}
