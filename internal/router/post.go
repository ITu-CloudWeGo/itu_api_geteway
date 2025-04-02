package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/post"
	"github.com/cloudwego/hertz/pkg/route"
)

func PostRoutes(r *route.RouterGroup) {
	r.POST("/post", post.SendPost)
	r.DELETE("/post", post.DeletePost)
	r.GET("/post", post.GetPost)
	r.PUT("/post", post.UpdatePost)
	r.GET("/post/user_posts", post.GetPostsByUid)
}
