package models

import (
	"git.querycap.com/practice/srv-demo-suns/pkg/constants/types"
	"github.com/go-courier/sqlx/v2"
	"time"
)

// 公共部分  db.go
var DB = sqlx.NewDatabase("mydb")

// 自增主键定义，原则上每个表都要有这个字段
type PrimaryID struct {
	// 自增 ID
	ID uint64 `db:"f_id,autoincrement" json:"-"`
}

// 操作时间定义，原则上每个表都要有这三个字段，有特殊要求可自行修改
// pgbuilder 在创建、修改时都需要 Mark 时间
type OperationTimes struct {
	// 创建时间
	CreatedAt types.Timestamp `db:"f_created_at,default='0'" json:"createdAt" `
	// 更新时间
	UpdatedAt types.Timestamp `db:"f_updated_at,default='0'" json:"updatedAt"`
}

func (times *OperationTimes) MarkUpdatedAt() {
	times.UpdatedAt = types.Timestamp(time.Now())
}

func (times *OperationTimes) MarkCreatedAt() {
	times.MarkUpdatedAt()
	times.CreatedAt = times.UpdatedAt
}

type OperationTimesWithDeletedAt struct {
	OperationTimes
	// 删除时间
	DeletedAt types.Timestamp `db:"f_deleted_at,default='0'" json:"-"`
}

func (times *OperationTimesWithDeletedAt) MarkDeletedAt() {
	times.MarkUpdatedAt()
	times.DeletedAt = times.UpdatedAt
}
