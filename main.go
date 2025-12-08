package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wonli/aqi"
)

func main() {
	//app是AppConfig结构体实例
	app := aqi.Init(
		//设置配置文件名
		aqi.ConfigFile("config-test.yaml"),
		//服务名称,以及服务端口
		aqi.HttpServer("my server", "port"),
	)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "serve running")
	})

	app.WithHttpServer(r)
	app.Start()
}
