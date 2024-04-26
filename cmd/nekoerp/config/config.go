package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	IsDebug       bool   `json:"isDebug"`       // 调试模式
	WebHost       string `json:"webHost"`       // 监听地址
	WebPort       int    `json:"webPort"`       // 监听端口
	MySQLHost     string `json:"mySQLHost"`     // MySQL主机
	MySQLPort     int    `json:"mySQLPort"`     // MySQL端口
	MySQLDataBase string `json:"mySQLDataBase"` // MySQL数据库
	MySQLUsername string `json:"mySQLUsername"` // MySQL用户名
	MySQLPassword string `json:"mySQLPassword"` // MySQL密码
	JWTScreenKet  string `json:"JWTScreenKet"`  // JWT密钥
}

func InitConfig() (*Config, error) {
	// 实例化Config对象
	var config Config

	// 读取配置文件
	configFile, err := os.ReadFile("config/config.json")
	if err != nil {
		return nil, err
	}

	// 将读取的文件内容解析到Config结构体中
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
