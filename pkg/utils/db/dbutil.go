package db

import (
	"strconv"

	"github.com/go-courier/sqlx/v2/builder"
)

func RightLikeOrIn(col *builder.Column, values []string) builder.SqlCondition {
	if len(values) == 1 {
		return col.RightLike(values[0])
	}
	return col.In(values)
}

func ConditionWhen(when bool, buildCondition func() builder.SqlCondition) builder.SqlCondition {
	if when {
		return buildCondition()
	}
	return nil
}

func ParseIDOrName(idOrName string) (uint64, string) {
	id, err := strconv.ParseUint(idOrName, 10, 64)
	if err != nil {
		return 0, idOrName
	}
	return id, ""
}
