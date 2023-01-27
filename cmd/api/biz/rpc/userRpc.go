package rpc

import (
	"github.com/cloudwego/kitex/client"
	"mini-tiktok-hanyongyan/cmd/user/kitex_gen/userservice/userservice"
	"mini-tiktok-hanyongyan/pkg/config"
)

var UserRpcClient userservice.Client

func UserRpcInit() {
	c, err := userservice.NewClient("UserService", client.WithHostPorts(config.UserServiceAddr))
	if err != nil {
		panic(err.Error())
	}
	UserRpcClient = c
}
