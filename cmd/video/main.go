package main

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/server"
	"log"
	"mini-tiktok-hanyongyan/cmd/video/kitex_gen/videoService/videoservice"
	"mini-tiktok-hanyongyan/cmd/video/rpc"
	"mini-tiktok-hanyongyan/pkg/config"
	"mini-tiktok-hanyongyan/pkg/dal/query"
	"net"
)

func Init() {
	// viper redis mysql 初始化
	config.Init()
}
func main() {
	// 设置端口
	addr, err := net.ResolveTCPAddr(config.TCP, config.VideoServiceAddr)

	// videoService 配置
	svr := videoservice.NewServer(
		new(VideoServiceImpl),
		server.WithServiceAddr(addr),
	)
	Init()
	// test
	user, _ := query.Q.TUser.WithContext(context.Background()).First()
	fmt.Println(user)
	fmt.Println(11111)
	// test end
	// 初始化 userRpcClient
	rpc.Init()
	// 启动 videoService
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
