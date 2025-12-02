package main

import "github.com/wonli/aqi"

func main() {
	aqi.Init(
		//自动解析工作目录，设置配置文件
		aqi.ConfigFile("config.yaml"),
	)

	/*
		支持传入Option函数
		type Option func(config *AppConfig) error


	*/
}
