package favorite

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/a"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_favorite/kitex_gen/favorite_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func AddFavoriteRoutes(ctx context.Context, c *app.RequestContext) {
	var req model.AddDelFavoriteRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	favoriteRpcCli := rpc_client.GetFavoriteRpcClient()
	resp, err := favoriteRpcCli.AddFavorite(ctx, &favorite_service.AddFavoriteRequest{
		Pid: req.Pid,
		Uid: req.Uid,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a.H{
		"status": http.StatusOK,
		"msg":    resp.Message,
	})
}
