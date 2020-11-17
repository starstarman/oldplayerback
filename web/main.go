package main

import (
	"oldplayerback/bootstrap"
	"oldplayerback/web/routes"
)

func newApp() *bootstrap.Bootstrapper {
	//初始化应用
	app := bootstrap.New("老玩家回归", "kuo")
	app.Bootstrap()
	app.Configure(routes.Configure)

	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
