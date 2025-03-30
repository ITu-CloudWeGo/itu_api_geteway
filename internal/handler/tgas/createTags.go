package tgas

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func CreateTags(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.CreateTagsRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
	}
	reqRPC, err := rpc_client.InitTagsRpcClient()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
		return
	}
	res, err := reqRPC.CreateTags(ctx, &tags_service.CreateTagsRequest{
		Pid:  req.Pid,
		Tags: req.Tags,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
		return
	}
	resp := model.CreateTagsResponse{
		Status: res.Status,
		Msg:    res.Msg,
	}
	c.JSON(200, resp)
}
