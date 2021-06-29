package global

import (
	"context"
	"git.querycap.com/tools/confpostgres"
	"git.querycap.com/tools/scaffold/pkg/appinfo"
	"git.querycap.com/tools/svcutil/confhttp"
	"git.querycap.com/tools/svcutil/conflogger"
	"github.com/go-courier/courier"
	"github.com/go-courier/sqlx/v2/migration"
	"srv-demo-suns/pkg/clients/client_id"
	"srv-demo-suns/pkg/models"
	"srv-demo-suns/pkg/utils/db"
	"srv-demo-suns/pkg/utils/idgen"
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
	postgres = &confpostgres.PostgresEndpoint{
		Database: models.DB,
	}

	idGen = idgen.MustNewIDGen()
)

var (
	// 添加一个上游服务 srv-id，ClientEndpoint负责对服务进行配置
	idClient = &confhttp.ClientEndpoint{}

	// 使用生成的NewClientxxx 方法实例化上游服务
	ClientID = client_id.NewClientID(idClient)
)

// 配置初始化
func init() {
	conflogger.SetProjectName(App.String())

	// 服务配置定义，需要配置的实体都需要在这里定义（若有数据库，上游服务，redis等也在这里定义）
	var Config = &struct {
		Log      *conflogger.Log
		Server   *confhttp.Server
		ClientID *confhttp.ClientEndpoint // 每个上游服务都需要在config 中添加，以便解析相关配置
		Postgres *confpostgres.PostgresEndpoint
	}{
		Server:   server,
		ClientID: idClient,
		Postgres: postgres,
	}

	// ConfP 方法会对项目的配置进行解析和配置，本地使用local.yml，集群中用master.yml
	App.ConfP(Config)

	confhttp.RegisterCheckerFromStruct(Config) // 注册调用链的健康检查
}

// 某些配置或全局实体可以通过 WithContextCompose 注入到服务的context中
var WithContext = confhttp.WithContextCompose(
	db.WithDBExecutor(postgres),
	idgen.WithIDGen(idGen),
)

func Server() courier.Transport {
	return server.WithContextInjector(WithContext)
}

func Migrate() {
	ctx, log := conflogger.NewContextAndLogger(context.Background(), "Migrate")
	defer log.End()

	if err := migration.Migrate(postgres.WithContext(ctx), nil); err != nil {
		panic(err)
	}
}
