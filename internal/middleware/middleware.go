package middleware

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service/authservice"
	"github.com/cloudwego/hertz/pkg/app"
	"log"
	"strings"
)

func CheckAccessToken(authClient authservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		tokenData := c.GetHeader("Authorization")
		accessToken := strings.TrimPrefix(string(tokenData), "Bearer ")
		req := &auth_service.CheckAccessTokenRequest{
			AccessToken: accessToken,
		}
		resp, err := authClient.CheckAccessToken(ctx, req)
		if err != nil {
			log.Printf("validate access token failed: %v", err)
			c.AbortWithStatus(401)
			log.Fatal(err)
		}
		if resp.CheckAccessExpiresAt == false {
			log.Printf("access token has not expired")
			c.AbortWithStatus(401)
			return
		}
		c.Next(ctx)
	}
}

func CheckRefreshToken(authClient authservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		tokenData := c.GetHeader("Authorization")
		refreshToken := strings.TrimPrefix(string(tokenData), "Bearer	")
		req := &auth_service.CheckRefreshTokenRequest{
			RefreshToken: refreshToken,
		}
		resp, err := authClient.CheckRefreshToken(ctx, req)
		if err != nil {
			log.Printf("validate refresh token failed: %v", err)
			c.AbortWithStatus(401)
			log.Fatal(err)
		}
		if resp.CheckRefreshExpiresAt == false {
			log.Printf("refresh token has not expired")
			c.AbortWithStatus(401)
			return
		}
		c.Next(ctx)
	}
}

func GenerateToken(authClient authservice.Client, ctx context.Context, uid uint64) (*auth_service.GenerateTokenResponse, error) {
	req := &auth_service.GenerateTokenRequest{
		Uid: uid,
	}
	resp, err := authClient.GenerateToken(ctx, req)
	if err != nil {
		log.Printf("generate token failed: %v", err)
		return nil, err
	}
	return resp, nil
}
