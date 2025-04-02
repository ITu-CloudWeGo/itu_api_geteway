package post

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/a"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_post/kitex_gen/post_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func DeletePost(ctx context.Context, c *app.RequestContext) {
	var req model.DeletePostReq
	if err := c.BindAndValidate(&req); err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	postRpcCli := rpc_client.GetPostRpcClient()
	resp, err := postRpcCli.DeletePost(ctx, &post_service.DeletePostRequest{
		PostId: req.Pid,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, a.H{
			"status": http.StatusInternalServerError,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a.H{
		"status": http.StatusOK,
		"msg":    resp.BaseResponse.Message,
	})
}
