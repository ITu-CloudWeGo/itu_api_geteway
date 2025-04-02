package post

import (
	"context"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/a"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_post/kitex_gen/post_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func GetPost(ctx context.Context, c *app.RequestContext) {
	pidStr := c.Query("pid")
	uidStr := c.Query("uid")
	if pidStr == "" {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    "pid is required",
		})
		return
	}
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    "pid is invalid",
		})
		return
	}
	var postResp *post_service.GetPostResponse
	postRpcCli := rpc_client.GetPostRpcClient()

	if uidStr != "" {
		uid, err := strconv.ParseInt(uidStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, a.H{
				"status": http.StatusBadRequest,
				"msg":    "pid is invalid",
			})
			return
		}
		postResp, err = postRpcCli.GetPost(ctx, &post_service.GetPostRequest{
			PostId: pid,
			UserId: uid,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, a.H{
				"status": http.StatusBadRequest,
				"msg":    err.Error(),
			})
			return
		}
	} else {
		postResp, err = postRpcCli.GetPostWithoutUid(ctx, &post_service.GetPostWithoutUidRequest{
			PostId: pid,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, a.H{
				"status": http.StatusBadRequest,
				"msg":    err.Error(),
			})
			return
		}
	}
	c.JSON(http.StatusOK, a.H{
		"status": http.StatusOK,
		"msg":    "success",
		"data":   postResp,
	})

}
