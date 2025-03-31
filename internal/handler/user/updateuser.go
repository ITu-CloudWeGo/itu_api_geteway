package user

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_user/kitex_gen/user_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func UpdateUser(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.UpdateUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
	}
	reqRPC, err := rpc_client.InitUserRpcClient()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
	}
	res, err := reqRPC.UpdateUser(ctx, &user_service.UpdateUserRequest{
		Uid:           req.Uid,
		NewEmail:      req.NewEmail,
		NewNickname:   req.NewNickname,
		NewIcon:       req.NewIcon,
		Password:      req.Password,
		NewPassword:   req.NewPassword,
		ConfirmPasswd: req.ConfirmPasswd,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
		return
	}
	resp := model.UpdateUserResponse{
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
