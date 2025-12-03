package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wonli/aqi"
)

func main() {
	app := aqi.Init(
		aqi.ConfigFile("config.yaml"),
		aqi.HttpServer("my server", "port"),
	)

	r := gin.Default()

	app.WithHttpServer(r)
	app.Start()
}
