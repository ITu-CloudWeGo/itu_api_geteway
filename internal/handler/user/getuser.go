package user

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_user/kitex_gen/user_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func GetUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.GetUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
	}
	reqRPC := rpc_client.GetUserRpcClient()
	//if err != nil {
	//	c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
	//	return
	//}
	res, err := reqRPC.GetUser(ctx, &user_service.GetUserRequest{
		Uid:      req.Uid,
		Username: req.UserName,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
		return
	}
	resp := model.GetUserResponse{
		Status: res.Response.Status,
		Msg:    res.Response.Msg,
		Data: model.Data{
			Username: res.Username,
			Nickname: res.Nickname,
			Icon:     res.Icon,
		},
	}
	c.JSON(200, resp)
}
