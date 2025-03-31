package rpc_client

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_api_geteway/config"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service/authservice"
	"github.com/ITu-CloudWeGo/itu_rpc_like/kitex_gen/likes_service/likesservice"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service/tagsservice"
	"github.com/ITu-CloudWeGo/itu_rpc_user/kitex_gen/user_service/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var (
	globalConfig *config.Config
)

func init() {
	var err error
	globalConfig, err = config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config:%v", err))
	}
}

// 认证模块
func InitAuthRpcClient() (cli authservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = authservice.NewClient(
		"auth",
		client.WithResolver(r),
	)
	if err != nil {
		return cli, err
	}
	return cli, err
}

// 标签模块
func InitTagsRpcClient() (cli tagsservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = tagsservice.NewClient(
		"tags",
		client.WithResolver(r),
	)
	if err != nil {
		return cli, err
	}
	return cli, err
}

//收藏模块

func InitLikesRpcClient() (cli likesservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = likesservice.NewClient(
		"likes",
		client.WithResolver(r),
	)
	if err != nil {
		return cli, err
	}
	return cli, err
}

//用户模块

func InitUserRpcClient() (cli userservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = userservice.NewClient(
		"user",
		client.WithResolver(r),
	)
	if err != nil {
		return cli, err
	}
	return cli, err
}
