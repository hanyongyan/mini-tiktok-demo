package config

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"mini-tiktok-hanyongyan/cmd/video/cos"
	"mini-tiktok-hanyongyan/pkg/cache"
	"mini-tiktok-hanyongyan/pkg/dal"
	"os"
)

func Init() {
	// 初始化 viper
	err := viperInit("./pkg/config/config.yaml")
	if err != nil {
		panic(err.Error())
	}
	// 初始化数据库
	dal.Init()
	// 初始化 Redis
	cache.Init()
	// 初始化 cos
	cos.Init()
}

func viperInit(path string) error {
	configFile, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	err = viper.ReadConfig(bytes.NewBuffer(configFile))
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	return nil
}
