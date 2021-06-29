package types

//go:generate tools gen enum UserType
type UserType uint8

const (
	USER_TYPE_UNKNOWN UserType = iota
	USER_TYPE__USER            // 普通用户
	USER_TYPE__ADMIN           // 管理员
)
