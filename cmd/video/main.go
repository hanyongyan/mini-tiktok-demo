package main

import (
	"github.com/cloudwego/kitex/server"
	"log"
	"mini-tiktok-hanyongyan/cmd/video/kitex_gen/videoService/videoservice"
	"mini-tiktok-hanyongyan/cmd/video/rpc"
	"mini-tiktok-hanyongyan/pkg/config"
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
	rpc.Init()
	// 启动 videoService
	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}

}
