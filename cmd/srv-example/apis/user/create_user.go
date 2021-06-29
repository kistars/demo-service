package user

import (
	"context"
	"git.querycap.com/practice/srv-demo-suns/pkg/controllers/user"
	"git.querycap.com/practice/srv-demo-suns/pkg/models"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
)

func init() {
	RouterUsers.Register(courier.NewRouter(&CreateUser{}))
}

type CreateUser struct {
	httpx.MethodPost `summary:"创建用户"`
	models.UserBase  `in:"body"`
}

func (req *CreateUser) Output(ctx context.Context) (interface{}, error) {
	return user.CreateUser(ctx, req.UserBase)
}
