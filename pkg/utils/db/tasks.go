package db

import (
	"context"

	"github.com/go-courier/sqlx/v2"
)

type Task func(ctx context.Context) error

type Tasks []Task

func (tasks Tasks) Do(ctx context.Context) (err error) {
	t := sqlx.NewTasks(DBExecutorFromContext(ctx))

	finalTasks := make([]sqlx.Task, len(tasks))

	for i := range tasks {
		finalTasks[i] = with(ctx, tasks[i])
	}

	return t.With(finalTasks...).Do()
}

func with(ctx context.Context, task Task) func(db sqlx.DBExecutor) error {
	return func(db sqlx.DBExecutor) error {
		return task(WithDBExecutor(db)(ctx))
	}
}
