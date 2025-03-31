package likes

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_like/kitex_gen/likes_service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func AddLikes(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.AddLikesRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}
	reqRPC := rpc_client.GetLikesRpcClient()
	//if err != nil {
	//	c.String(consts.StatusInternalServerError, err.Error())
	//	return
	//}
	res, err := reqRPC.AddLikes(ctx, &likes_service.AddLikesRequest{
		Pid: req.Pid,
		Uid: req.Uid,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp := model.AddLikesResponse{
		Status: res.Status,
		Msg:    res.Msg,
	}
	c.JSON(200, resp)
}
