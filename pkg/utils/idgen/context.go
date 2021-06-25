package idgen

import (
	"context"
)

type contextKeyIDGenerator int

func WithIDGen(idGen IDGen) func(ctx context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, contextKeyIDGenerator(0), idGen)
	}
}

func IDGenFromContext(ctx context.Context) IDGen {
	return ctx.Value(contextKeyIDGenerator(0)).(IDGen)
}
