package rpc

import (
	"github.com/cloudwego/kitex/client"
	"mini-tiktok-hanyongyan/cmd/video/kitex_gen/videoService/videoservice"
	"mini-tiktok-hanyongyan/pkg/config"
)

var VideoRpcClient videoservice.Client

func VideoRpcInit() {
	c, err := videoservice.NewClient("videoService", client.WithHostPorts(config.VideoServiceAddr))
	if err != nil {
		panic(err.Error())
	}
	VideoRpcClient = c
}
