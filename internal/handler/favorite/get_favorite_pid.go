package favorite

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/a"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_favorite/kitex_gen/favorite_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func GetFavoritePid(ctx context.Context, c *app.RequestContext) {
	uidStr := c.Query("uid")
	if uidStr == "" {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    "uid is required",
		})
		return
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    "uid is invalid",
		})
		return
	}
	favoriteRpcCli := rpc_client.GetFavoriteRpcClient()
	resp, err := favoriteRpcCli.GetFavoritePostsIdByUid(ctx, &favorite_service.GetFavoritePostsIdByUidReq{
		Uid: uid,
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
		"msg":    "success",
		"data":   resp.Pid,
	})
}
