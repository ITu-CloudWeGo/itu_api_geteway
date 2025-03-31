package likes

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_like/kitex_gen/likes_service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func DelLikes(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.DelLikesRequest
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
	res, err := reqRPC.DelLikes(ctx, &likes_service.DelLikesRequest{
		Uid: req.Uid,
		Pid: req.Pid,
	})
	if err != nil {
		c.String(consts.StatusInternalServerError, err.Error())
		return
	}
	resp := model.DelLikesResponse{
		Status: res.Status,
		Msg:    res.Msg,
	}
	c.JSON(200, resp)
}
