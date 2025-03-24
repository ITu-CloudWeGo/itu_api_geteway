package main

import (
	"fmt"
	"github.com/ITu-CloudWeGo/itu_api_geteway/config"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	conf, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	h := server.New(server.WithHostPorts(fmt.Sprintf("%s:%d", conf.Hertz.Host, conf.Hertz.Port)))
	h.Spin()
}
