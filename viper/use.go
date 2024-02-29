package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	// 配置文件设置
	viper.SetConfigName("config")        // 配置文件名称 (不带后缀)
	viper.SetConfigType("toml")          // 如果配置文件的名称不是通过 SetConfigName 设置，则需要设置配置文件类型
	viper.AddConfigPath("./viper/conf/") // 查找配置文件的路径

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("错误: %s\n", err)
	}

	// 读取配置项
	host := viper.GetString("server.host") // 读取 server.host 的值
	port := viper.GetInt("server.port")    // 读取 server.port 的值

	fmt.Printf("服务器地址: %s, 端口: %d\n", host, port)
}
