package subscribe

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/a"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_subscribe/kitex_gen/subscribe_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func DelSubscribe(ctx context.Context, c *app.RequestContext) {
	var req model.AddDelSubscribeRequest
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	subscribeRpcCli := rpc_client.GetSubscribeRpcClient()
	resp, err := subscribeRpcCli.RemoveSubscribe(ctx, &subscribe_service.RemoveSubscribeRequest{
		Uid:       req.Uid,
		AuthorUid: req.AuthorId,
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
