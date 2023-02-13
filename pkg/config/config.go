package config

// Author: Hanyongyan

// 用于存储配置信息
type GlobalConfig struct {
	DBConfig struct {
		Addr     string
		Username string
		Password string
		Database string
	} `mapstructure:"db"`
	RedisConfig struct {
		Addr     string
		Password string
		Database int
	} `mapstructure:"redis"`
	CosConfig struct {
		Secretid  string
		Secretkey string
		Url       string
		VideoPath string
		PhotoPath string
	} `mapstructure:"cos"`
}
