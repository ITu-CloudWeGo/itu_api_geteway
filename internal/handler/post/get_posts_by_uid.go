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

func GetPostsByUid(ctx context.Context, c *app.RequestContext) {
	uidStr := c.Query("uid")
	pageStr := c.Query("page")
	pageSizeStr := c.Query("page_size")
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	pageSize, err := strconv.ParseInt(pageSizeStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	postRpcCli := rpc_client.GetPostRpcClient()
	postResp, err := postRpcCli.GetUserPosts(ctx, &post_service.GetUserPostsRequest{
		UserId:   uid,
		Page:     int32(page),
		PageSize: int32(pageSize),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, a.H{
			"status": http.StatusInternalServerError,
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, a.H{
		"status": http.StatusOK,
		"msg":    "success",
		"data":   postResp.Posts,
	})
}
