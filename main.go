package main

import (
	"gin-web-demo/bootstrap"
	"gin-web-demo/web/router"

	"fmt"
	pldconf "gin-web-demo/config"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("gin-web-demo", "gin-web-demo")
	app.Bootstrap()
	app.Configure(router.Configure)
	return app
}

func main() {
	app := newApp()
	//读取配置文件，获取监听端口
	pldConf := pldconf.AppConfig
	port := pldConf.Server.Port
	listenPort := fmt.Sprintf(":%v", port) //格式化监听端口
	app.Listen(listenPort)
}
