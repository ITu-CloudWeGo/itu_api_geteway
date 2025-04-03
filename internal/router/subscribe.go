package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/subscribe"
	"github.com/cloudwego/hertz/pkg/route"
)

func SubscribeRoutes(r *route.RouterGroup) {
	r.POST("/subscribe", subscribe.AddSubscribe)
	r.DELETE("/subscribe", subscribe.DelSubscribe)
	r.GET("/subscribe/list", subscribe.GetSubsAuthorIds)
}
