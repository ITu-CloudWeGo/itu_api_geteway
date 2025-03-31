package post

import (
	"context"
	"encoding/json"
	"github.com/ITu-CloudWeGo/itu_api_geteway/internal/a"
	"github.com/ITu-CloudWeGo/itu_api_geteway/rpc_client"
	"github.com/ITu-CloudWeGo/itu_rpc_file/kitex_gen/file_service"
	"github.com/ITu-CloudWeGo/itu_rpc_post/kitex_gen/post_service"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
	"strconv"
)

func SendPost(ctx context.Context, c *app.RequestContext) {
	// TODO: Implement this function
	formData, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}
	postRpcCli := rpc_client.GetPostRpcClient()
	fileRpcCLi := rpc_client.GetFileRpcClient()
	tagRpcCli := rpc_client.GetTagsRpcClient()

	title := formData.Value["title"][0]
	content := formData.Value["content"][0]
	tags := formData.Value["tags"]
	imageFiles := formData.File["files"]
	uidStr := formData.Value["uid"][0]

	var imgUrls []*file_service.UploadFileResponse

	// 上传文件
	for _, imageFile := range imageFiles {
		file, err := imageFile.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, a.H{
				"status": http.StatusBadRequest,
				"msg":    err.Error(),
			})
			return
		}
		var fileContent []byte
		_, err = file.Read(fileContent)
		if err != nil {
			c.JSON(http.StatusBadRequest, a.H{
				"status": http.StatusBadRequest,
				"msg":    err.Error(),
			})
			return
		}
		fileResp, err := fileRpcCLi.UploadFile(ctx, &file_service.UploadFileRequest{
			FileName:    imageFile.Filename,
			FileContent: fileContent,
			Type:        PostFileType,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, a.H{
				"status": http.StatusBadRequest,
				"msg":    err.Error(),
			})
		}
		imgUrls = append(imgUrls, fileResp)
	}
	uid, _ := strconv.ParseInt(uidStr, 10, 64)
	imgUrlJson, err := json.Marshal(imgUrls)
	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}

	// 发送帖子
	postResp, err := postRpcCli.SendPost(ctx, &post_service.SendPostRequest{
		Title:   title,
		Content: content,
		UserId:  uid,
		Images:  string(imgUrlJson),
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, a.H{
			"status": http.StatusBadRequest,
			"msg":    err.Error(),
		})
		return
	}

	// 上传标签
	for _, tag := range tags {
		_, err := tagRpcCli.PidTidCreate(ctx, &tags_service.PidTidCreateRequest{
			Tag: tag,
			Pid: postResp.Pid,
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
		"msg":    postResp.BaseResponse.Message,
	})
}
