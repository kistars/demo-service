package main

import (
	"git.querycap.com/practice/srv-demo-suns/cmd/srv-example/apis"
	"git.querycap.com/practice/srv-demo-suns/cmd/srv-example/global"
	"github.com/go-courier/courier"
)

func main() {
	global.App.AddCommand("migrate", func(args ...string) {
		global.Migrate()
	})
	// 传入路由，运行服务
	global.App.Execute(func(args ...string) {
		courier.Run(apis.RouterRoot, global.Server())
	})
}
