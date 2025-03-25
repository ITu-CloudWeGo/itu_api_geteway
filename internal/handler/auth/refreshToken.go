package auth

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/module"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	var err error
	var req module.RefreshTokenRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	reqRpc, err := rpc_client.InitAuthRpcClient()
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client:"+err.Error())
		return
	}
	res, err := reqRpc.RefreshToken(ctx, &auth_service.RefreshTokenRequest{
		RefreshToken: req.RefreshToken,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to get Rpc client")
		return
	}
	resp := module.RefreshTokenResponse{
		AccessToken:        res.AccessToken,
		AccessTokenExpire:  res.AccessTokenExpire,
		RefreshToken:       res.RefreshToken,
		RefreshTokenExpire: res.RefreshTokenExpire,
	}
	c.JSON(200, resp)
}
