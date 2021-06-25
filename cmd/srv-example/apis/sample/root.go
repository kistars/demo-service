package sample

import (
	"github.com/go-courier/courier"
	"github.com/go-courier/httptransport"
)

var RootRouter = courier.NewRouter(httptransport.Group("/Id"))
