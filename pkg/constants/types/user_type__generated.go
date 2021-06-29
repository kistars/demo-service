package types

import (
	bytes "bytes"
	database_sql_driver "database/sql/driver"
	errors "errors"

	github_com_go_courier_enumeration "github.com/go-courier/enumeration"
)

var InvalidUserType = errors.New("invalid UserType type")

func ParseUserTypeFromLabelString(s string) (UserType, error) {
	switch s {
	case "":
		return USER_TYPE_UNKNOWN, nil
	case "普通用户":
		return USER_TYPE__USER, nil
	case "管理员":
		return USER_TYPE__ADMIN, nil
	}
	return USER_TYPE_UNKNOWN, InvalidUserType
}

func (v UserType) String() string {
	switch v {
	case USER_TYPE_UNKNOWN:
		return ""
	case USER_TYPE__USER:
		return "USER"
	case USER_TYPE__ADMIN:
		return "ADMIN"
	}
	return "UNKNOWN"
}

func ParseUserTypeFromString(s string) (UserType, error) {
	switch s {
	case "":
		return USER_TYPE_UNKNOWN, nil
	case "USER":
		return USER_TYPE__USER, nil
	case "ADMIN":
		return USER_TYPE__ADMIN, nil
	}
	return USER_TYPE_UNKNOWN, InvalidUserType
}

func (v UserType) Label() string {
	switch v {
	case USER_TYPE_UNKNOWN:
		return ""
	case USER_TYPE__USER:
		return "普通用户"
	case USER_TYPE__ADMIN:
		return "管理员"
	}
	return "UNKNOWN"
}

func (v UserType) Int() int {
	return int(v)
}

func (UserType) TypeName() string {
	return "srv-demo-suns/pkg/constants/types.UserType"
}

func (UserType) ConstValues() []github_com_go_courier_enumeration.IntStringerEnum {
	return []github_com_go_courier_enumeration.IntStringerEnum{USER_TYPE__USER, USER_TYPE__ADMIN}
}

func (v UserType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidUserType
	}
	return []byte(str), nil
}

func (v *UserType) UnmarshalText(data []byte) (err error) {
	*v, err = ParseUserTypeFromString(string(bytes.ToUpper(data)))
	return
}

func (v UserType) Value() (database_sql_driver.Value, error) {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}
	return int64(v) + int64(offset), nil
}

func (v *UserType) Scan(src interface{}) error {
	offset := 0
	if o, ok := (interface{})(v).(github_com_go_courier_enumeration.DriverValueOffset); ok {
		offset = o.Offset()
	}

	i, err := github_com_go_courier_enumeration.ScanIntEnumStringer(src, offset)
	if err != nil {
		return err
	}
	*v = UserType(i)
	return nil
}
