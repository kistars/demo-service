package apis

import (
	"git.querycap.com/tools/svcutil/confhttp"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
	"github.com/go-courier/httptransport/openapi"
)

// 服务根路径（er / Liveness 需注册在跟路径下）
var RouterRoot = courier.NewRouter(httptransport.Group("/"))

// 服务根路径后必须是服务名（小写服务名且不包含 srv- 前缀）
var RouterExample = courier.NewRouter(httptransport.Group("/example"))

func init() {
	// 注册服务探活
	RouterRoot.Register(confhttp.LivenessRouter)
	// 注册服务名路径
	RouterRoot.Register(RouterExample)
	// 注册OpenAPI
	RouterExample.Register(openapi.OpenAPIRouter)
}
