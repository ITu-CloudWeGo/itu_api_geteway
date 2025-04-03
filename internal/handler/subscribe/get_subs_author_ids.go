package subscribe

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/a"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_subscribe/kitex_gen/subscribe_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func GetSubsAuthorIds(ctx context.Context, c *app.RequestContext) {
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
	subsRpcCli := rpc_client.GetSubscribeRpcClient()
	subsResp, err := subsRpcCli.GetSubscribe(ctx, &subscribe_service.GetSubscribeReq{
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
		"msg":    "get subs author ids successfully",
		"data":   subsResp.AuthorId,
	})
}
