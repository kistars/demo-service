package global

import (
	"git.querycap.com/tools/confpostgres"
	"git.querycap.com/tools/scaffold/pkg/appinfo"
	"git.querycap.com/tools/svcutil/confhttp"
	"git.querycap.com/tools/svcutil/conflogger"
	"github.com/go-courier/courier"
	"srv-demo-suns/pkg/models"
)

var (
	server = &confhttp.Server{} // 可配置的Server实体，实现了Serve接口

	App = appinfo.New( // App代表了这个服务，包含了各种信息，能解析服务的配置
		// 这里填服务名的全称 srv-xxx
		appinfo.WithName("srv-example"),
		appinfo.WithMainRoot(".."),
	)
)

var (
	// 数据库实例配置
	Postgres = &confpostgres.PostgresEndpoint{
		Database: models.DB,
	}
)

// 配置初始化
func init() {
	// 服务配置定义，需要配置的实体都需要在这里定义（若有数据库，上游服务，redis等也在这里定义）
	var Config = &struct {
		Log      *conflogger.Log
		Server   *confhttp.Server
		Postgres *confpostgres.PostgresEndpoint
	}{
		Server:   server,
		Postgres: Postgres,
	}

	// ConfP 方法会对项目的配置进行解析和配置，本地使用local.yml，集群中用master.yml
	App.ConfP(Config)

	confhttp.RegisterCheckerFromStruct(Config) // 注册调用链的健康检查
}

// 某些配置或全局实体可以通过 WithContextCompose 注入到服务的context中
var WithContext = confhttp.WithContextCompose()

func Server() courier.Transport {
	return server.WithContextInjector(WithContext)
}
