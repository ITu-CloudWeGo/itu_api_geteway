package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/likes"
	"github.com/cloudwego/hertz/pkg/route"
)

func LikesRoutes(r *route.RouterGroup) {
	r.POST("/like/likes", likes.AddLikes)
	r.DELETE("/like/unlikes", likes.DelLikes)
}
