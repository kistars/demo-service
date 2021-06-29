package user

import (
	"context"
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport/httpx"
)

func init() {
	RouterUsers.Register(courier.NewRouter(&ListUser{}))
}

// 用户列表 （--这里也是接口注释）
// 以 summary 中定义的为准
type ListUser struct {
	httpx.MethodGet `summary:"用户列表"` // summary tag为接口注释
}

type User struct {
	// 用户ID
	UserID int64 `json:"userID"`
	// 姓名
	Name string `json:"name"`
}

type UserDataList struct {
	Data  []User `json:"data"`
	Total int    `json:"total"`
}

// courier框架中的最小处理单元，operator
func (req *ListUser) Output(ctx context.Context) (interface{}, error) {
	list := make([]User, 0)

	list = append(list, User{
		Name:   "蜗牛",
		UserID: 10000,
	})

	return &UserDataList{
		Data:  list,
		Total: len(list),
	}, nil
}
