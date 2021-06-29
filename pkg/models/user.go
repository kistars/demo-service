package models

import "srv-demo-suns/pkg/constants/types"

//go:generate tools gen model2 User --database=DB
// 用户表
// @def primary ID
// @def unique_index I_user_id UserID
// @def unique_index I_phone Phone
// @def index I_name Name
// @def index i_account_type UserType
type User struct {
	PrimaryID
	RefUser
	UserBase
	OperationTimesWithDeletedAt
}

type UserBase struct {
	// 名称
	Name string `db:"f_name,size=50" json:"name"`
	// 手机号
	Phone string `db:"f_phone,size=11,default=''" json:"phone"`
	// 用户类型
	UserType types.UserType `db:"f_user_type" json:"userType"`
}

type UserID = types.SFID

type RefUser struct {
	// @rel User.UserID
	// 用户唯一 ID
	UserID UserID `db:"f_user_id" json:"userID"`
}
