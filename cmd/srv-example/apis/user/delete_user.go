package user

import (
	"context"
	"git.querycap.com/practice/srv-demo-suns/pkg/constants/types"
	"git.querycap.com/practice/srv-demo-suns/pkg/controllers/user"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
)

func init() {
	RouterUsers.Register(courier.NewRouter(&DeleteUser{}))
}

type DeleteUser struct {
	httpx.MethodDelete `summary:"删除用户" path:"/:userID"`
	// 用户 ID
	UserID types.SFID `name:"userID" in:"path"`
}

func (d *DeleteUser) Output(ctx context.Context) (interface{}, error) {
	return nil, user.DeleteUser(ctx, d.UserID)
}
