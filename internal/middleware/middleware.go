package middleware

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strings"
)

func CheckAccessToken() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		tokenData := c.GetHeader("Authorization")

		accessToken := strings.TrimPrefix(string(tokenData), "Bearer ")
		req := &auth_service.CheckAccessTokenRequest{
			AccessToken: accessToken,
		}
		authClient, err := rpc_client.InitAuthRpcClient()
		if err != nil {
			// Todo Log
			c.AbortWithStatus(401)
			return
		}
		resp, err := authClient.CheckAccessToken(ctx, req)
		if err != nil {
			log.Printf("validate access token failed: %v", err)
			c.AbortWithStatus(401)
			log.Fatal(err)
			return
		}
		if resp.CheckAccessExpiresAt == false {
			log.Printf("access token has not expired")
			c.AbortWithStatus(401)
			return
		}
		c.Next(ctx)
	}
}

func CheckRefreshToken() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		tokenData := c.GetHeader("Authorization")
		refreshToken := strings.TrimPrefix(string(tokenData), "Bearer	")
		req := &auth_service.CheckRefreshTokenRequest{
			RefreshToken: refreshToken,
		}
		authClient, err := rpc_client.InitAuthRpcClient()
		if err != nil {
			// Todo Log
			c.AbortWithStatus(401)
			return
		}
		resp, err := authClient.CheckRefreshToken(ctx, req)
		if err != nil {
			log.Printf("validate refresh token failed: %v", err)
			c.AbortWithStatus(401)
			log.Fatal(err)
			return
		}
		if resp.CheckRefreshExpiresAt == false {
			log.Printf("refresh token has not expired")
			c.AbortWithStatus(401)
			return
		}
		c.Next(ctx)
	}
}
