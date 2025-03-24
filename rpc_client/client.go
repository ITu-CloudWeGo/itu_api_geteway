package rpc_client

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_api_geteway/config"
	"github.com/ITu-CloudWeGo/itu_rpc_auth/rpc/kitex_gen/auth_service/authservice"
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
	r, err := etcd.NewEtcdResolver([]string{globalConfig.Registry.RegisterAddress})
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
