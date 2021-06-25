package db

import (
	"context"

	"github.com/go-courier/sqlx/v2"
)

type contextKeyDBExecutor int

func WithDBExecutor(db sqlx.DBExecutor) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, contextKeyDBExecutor(0), db)
	}
}

func DBExecutorFromContext(ctx context.Context) sqlx.DBExecutor {
	return ctx.Value(contextKeyDBExecutor(0)).(sqlx.DBExecutor).WithContext(ctx)
}
