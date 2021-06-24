package user

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

// users是业务根路径
var RouterUsers = courier.NewRouter(httptransport.Group("/users"))
