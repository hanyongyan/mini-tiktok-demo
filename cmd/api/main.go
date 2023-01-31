// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"mini-tiktok-hanyongyan/cmd/api/biz/rpc"
	"mini-tiktok-hanyongyan/pkg/config"
)

func Init() {
	//  viper mysql redis init
	config.Init()
	// userService and videoService init
	rpc.Init()
}

func main() {
	Init()
	h := server.Default(server.WithHostPorts(":8080"))

	register(h)

	h.Spin()
}
