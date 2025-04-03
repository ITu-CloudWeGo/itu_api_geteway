package router

import (
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/handler/favorite"
	"github.com/cloudwego/hertz/pkg/route"
)

func FavoriteRoutes(r *route.RouterGroup) {
	r.POST("/favorite", favorite.AddFavoriteRoutes)
	r.DELETE("/favorite", favorite.DelFavorite)
	r.GET("/favorite/list", favorite.GetFavoritePid)
}
