package rpc

import (
	"github.com/cloudwego/kitex/client"
	"mini-tiktok-hanyongyan/cmd/user/kitex_gen/userservice/userservice"
	"mini-tiktok-hanyongyan/pkg/config"
)

var UserRpcClient userservice.Client

// Init userService rpc init
func Init() {
	c, err := userservice.NewClient("userService", client.WithHostPorts(config.UserServiceAddr))
	if err != nil {
		panic(err.Error())
	}
	UserRpcClient = c
}
