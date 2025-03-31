package rpc_client

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_api_geteway/config"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service/authservice"
	"github.com/ITu-CloudWeGo/itu_rpc_favorite/kitex_gen/favorite_service/favoriteservice"
	"github.com/ITu-CloudWeGo/itu_rpc_file/kitex_gen/file_service/fileservice"
	"github.com/ITu-CloudWeGo/itu_rpc_like/kitex_gen/likes_service/likesservice"
	"github.com/ITu-CloudWeGo/itu_rpc_post/kitex_gen/post_service/postservice"
	"github.com/ITu-CloudWeGo/itu_rpc_subscribe/kitex_gen/subscribe_service/subscribeservice"
	"github.com/ITu-CloudWeGo/itu_rpc_tags/kitex_gen/tags_service/tagsservice"
	"github.com/ITu-CloudWeGo/itu_rpc_user/kitex_gen/user_service/userservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var (
	globalConfig       *config.Config
	authRpcClient      authservice.Client
	tagsRpcClient      tagsservice.Client
	likesRpcClient     likesservice.Client
	userRpcClient      userservice.Client
	postRpcClient      postservice.Client
	favoriteRpcClient  favoriteservice.Client
	fileRpcClient      fileservice.Client
	subscribeRpcClient subscribeservice.Client
)

func init() {
	var err error
	globalConfig, err = config.GetConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to load config:%v", err))
	}
	authRpcClient, err = initAuthRpcClient()
	if err != nil {
		panic(err)
	}
	tagsRpcClient, err = initTagsRpcClient()
	if err != nil {
		panic(err)
	}
	likesRpcClient, err = initLikesRpcClient()
	if err != nil {
		panic(err)
	}
	userRpcClient, err = initUserRpcClient()
	if err != nil {
		panic(err)
	}
	postRpcClient, err = initPostRpcClient()
	if err != nil {
		panic(err)
	}

}

// 认证模块
func initAuthRpcClient() (cli authservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = authservice.NewClient(
		"auth_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetAuthRpcClient() (cli authservice.Client) {
	return authRpcClient
}

// 标签模块
func initTagsRpcClient() (cli tagsservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = tagsservice.NewClient(
		"tag_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetTagsRpcClient() (cli tagsservice.Client) {
	return tagsRpcClient
}

// 收藏模块
func initLikesRpcClient() (cli likesservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = likesservice.NewClient(
		"like_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetLikesRpcClient() (cli likesservice.Client) {
	return likesRpcClient
}

// 用户模块
func initUserRpcClient() (cli userservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = userservice.NewClient(
		"user_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetUserRpcClient() (cli userservice.Client) {
	return userRpcClient
}

// 帖子模块
func initPostRpcClient() (cli postservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = postservice.NewClient(
		"post_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetPostRpcClient() (cli postservice.Client) {
	return postRpcClient
}

// 收藏模块
func initFavoriteRpcClient() (cli favoriteservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = favoriteservice.NewClient(
		"favorite_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetFavoriteRpcClient() (cli favoriteservice.Client) {
	return favoriteRpcClient
}

// 文件模块
func initFileRpcClient() (cli fileservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = fileservice.NewClient(
		"file_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetFileRpcClient() (cli fileservice.Client) {
	return fileRpcClient
}

// 订阅模块
func initSubscribeRpcClient() (cli subscribeservice.Client, err error) {
	r, err := etcd.NewEtcdResolver(globalConfig.Registry.RegisterAddress)
	if err != nil {
		log.Fatal(err)
	}
	cli, err = subscribeservice.NewClient(
		"subscribe_service",
		client.WithResolver(r),
	)
	if err != nil {
		return nil, err
	}
	return cli, err
}
func GetSubscribeRpcClient() (cli subscribeservice.Client) {
	return subscribeRpcClient
}
