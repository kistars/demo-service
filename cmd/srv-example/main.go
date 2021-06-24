package main

import (
	"srv-demo-suns/cmd/srv-example/apis"
	"srv-demo-suns/cmd/srv-example/global"

	"github.com/go-courier/courier"
)

func main() {
	// 传入路由，运行服务
	global.App.Execute(func(args ...string) {
		courier.Run(apis.RouterRoot, global.Server())
	})
}
