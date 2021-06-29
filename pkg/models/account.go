package models

import "git.querycap.com/practice/srv-demo-suns/pkg/constants/types"

//go:generate tools gen model2 Account  --database=DB
// 账户 --表注释
// @def 部分：定义表的索引，格式为：@def 索引类型 索引名 字段名
// @def primary ID
// @def unique_index I_account_id AccountID
// @def index I_nick_name NickName
// @def index I_created_at CreatedAt
type Account struct {
	PrimaryID
	RefAccount
	AccountBase

	// 状态 --字段注释
	State int `db:"f_state,default='1'" json:"state"`

	OperationTimes
}

type AccountBase struct {
	// 昵称
	NickName string `db:"f_nick_name,size=20" json:"nickName"`
	// 账户类型
	AccountType int `db:"f_account_type" json:"accountType"`
	// 头像
	Avatar string `db:"f_avatar,default=''" json:"avatar"`
}

// Ref结构体表明该字段有其他表也需要用到，通过‘@rel 表名.字段名’来表示关联
type RefAccount struct {
	// @rel Account.AccountID
	// 账户唯一 ID
	AccountID types.SFID `db:"f_account_id" json:"accountID"`
}
