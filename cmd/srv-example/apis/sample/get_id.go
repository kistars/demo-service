package sample

import (
	"context"
	"git.querycap.com/practice/srv-demo-suns/cmd/srv-example/global"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
)

func init() {
	RootRouter.Register(courier.NewRouter(GetID{}))
}

// 获取 ID
type GetID struct {
	httpx.MethodGet `summary:"获取 ID"`
}

func (req GetID) Path() string {
	return ""
}

func (req GetID) Output(ctx context.Context) (interface{}, error) {
	// 最简单的方式即直接调用ClientID的对应方法，若有其他逻辑，则建议放到modules下面，参见idp
	resp, _, err := global.ClientID.WithContext(ctx).GenerateID()
	if err != nil {
		return nil, err
	}
	return resp.ID, nil
}
