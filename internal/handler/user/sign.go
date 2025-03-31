package user

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_user/kitex_gen/user_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func Sign(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.SignRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
	}
	reqRPC := rpc_client.GetUserRpcClient()
	//if err != nil {
	//	c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
	//	return
	//}
	res, err := reqRPC.Sign(ctx, &user_service.SignRequest{
		Email:    req.Email,
		Username: req.Username,
		Password: req.Password,
		Nickname: req.Nickname,
		Icon:     req.Email,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
		return
	}
	resp := model.SignResponse{
		Status: res.Response.Status,
		Msg:    res.Response.Msg,
		Data: model.Data{
			Email:    res.Email,
			Username: res.Username,
			Nickname: res.Nickname,
			Icon:     res.Icon,
		},
	}
	c.JSON(200, resp)
}
