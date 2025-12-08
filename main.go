package main

import (
	"net/http"

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

	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("serve running"))
	})

	app.WithHttpServer(h)
	app.Start()
}
