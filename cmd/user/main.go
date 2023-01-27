package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"mini-tiktok-hanyongyan/cmd/user/kitex_gen/userservice/userservice"
	"mini-tiktok-hanyongyan/cmd/user/rpc"
	"mini-tiktok-hanyongyan/pkg/config"
	"net"
)

func Init() {
	// viper redis mysql 初始化
	config.Init()
}

func main() {
	addr, err := net.ResolveTCPAddr(config.TCP, config.UserServiceAddr)
	// userService 的配置
	svr := userservice.NewServer(
		new(UserServiceImpl),
		// 指定端口启动
		server.WithServiceAddr(addr),
	)
	Init()
	// videoService rpc init
	rpc.Init()
	// 启动 userService
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
