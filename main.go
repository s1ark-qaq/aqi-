package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wonli/aqi"
)

func main() {
	app := aqi.Init(
		//CommitVersion为空,也就是开发版本时,配置文件都会加上-dev后缀,在项目开始时生成yaml配置文件,其中有部分默认的配置.
		aqi.ConfigFile("config-test.yaml"),
		//服务名称,以及服务端口,默认2015
		aqi.HttpServer("my server", "port"),
	)

	r := gin.Default()

	app.WithHttpServer(r)
	app.Start()
}
