package user

import (
	"context"
	"git.querycap.com/practice/srv-demo-suns/pkg/constants/types"
	"git.querycap.com/practice/srv-demo-suns/pkg/controllers/user"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
)

func init() {
	RouterUsers.Register(courier.NewRouter(&GetUser{}))
}

type GetUser struct {
	httpx.MethodGet `summary:"获取用户详情" path:"/:userID"`
	UserID          types.SFID `name:"userID" in:"path"`
}

func (req *GetUser) Output(ctx context.Context) (interface{}, error) {
	return user.RetrieveUser(ctx, req.UserID)
}
