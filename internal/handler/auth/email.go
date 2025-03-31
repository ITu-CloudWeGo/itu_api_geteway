package auth

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/model"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service"
	"github.com/cloudwego/hertz/pkg/app"
)

func EmailVerify(ctx context.Context, c *app.RequestContext) {
	var err error
	var req model.EmailRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(500, "Failed to get Rpc client:"+err.Error())
	}
	reqRPC := rpc_client.GetAuthRpcClient()
	//if err != nil {
	//	c.String(500, "Failed to get Rpc client:"+err.Error())
	//	return
	//}
	res, err := reqRPC.Email(ctx, &auth_service.EmailRequest{
		Email: req.Email,
	})
	if err != nil {
		c.String(500, "Failed to get Rpc client:"+err.Error())
		return
	}
	resp := model.EmailResponse{
		Captcha: res.Captcha,
	}
	c.JSON(200, resp)
}
