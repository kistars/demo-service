package user

import (
	"context"
	"git.querycap.com/practice/srv-demo-suns/pkg/constants/types"
	"git.querycap.com/practice/srv-demo-suns/pkg/controllers/user"
	"git.querycap.com/practice/srv-demo-suns/pkg/models"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
)

func init() {
	RouterUsers.Register(courier.NewRouter(&UpdateUser{}))
}

type UpdateUser struct {
	httpx.MethodPut `summary:"修改用户" path:"/:userID"`
	// 用户 ID
	UserID          types.SFID `name:"userID" in:"path"`
	models.UserBase `in:"body"`
}

func (req *UpdateUser) Output(ctx context.Context) (interface{}, error) {
	return nil, user.UpdateUser(ctx, req.UserBase, req.UserID)
}
