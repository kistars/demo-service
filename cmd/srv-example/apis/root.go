package apis

import (
	"git.querycap.com/tools/svcutil/confhttp"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
	"github.com/go-courier/httptransport/openapi"
	"srv-demo-suns/cmd/srv-example/apis/sample"
	"srv-demo-suns/cmd/srv-example/apis/user"
)

// 服务根路径（er / Liveness 需注册在跟路径下）
var RouterRoot = courier.NewRouter(httptransport.Group("/"))

// 服务根路径后必须是服务名（小写服务名且不包含 srv- 前缀）
var RouterExample = courier.NewRouter(httptransport.Group("/example"))

// 版本号（这里的v0）必须跟在服务名后面，因此每个服务的前两级服务都是 /xxx/v0 格式
var RouterV0 = courier.NewRouter(httptransport.Group("/v0"))

func init() {
	// 注册服务探活
	RouterRoot.Register(confhttp.LivenessRouter)
	// 注册服务名路径
	RouterRoot.Register(RouterExample)
	// 注册OpenAPI
	RouterExample.Register(openapi.OpenAPIRouter)
	// 注册v0路径
	RouterExample.Register(RouterV0)

	{
		// 在v0下注册真正的业务路径（也就是前面创建的user根路径）
		RouterV0.Register(user.RouterUsers)
		// getId路由
		RouterV0.Register(sample.RootRouter)
	}
}
