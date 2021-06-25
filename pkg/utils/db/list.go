package db

import (
	"context"
	"github.com/go-courier/sqlx-pg/pgbuilder"
	"github.com/go-courier/sqlx/v2"
	"github.com/go-courier/sqlx/v2/builder"
)

type AdditionsBuilder interface {
	ToAdditions(db sqlx.DBExecutor) builder.Additions
}

type Actions = []Action
type Action = func(ctx context.Context) error

type WithPostActions interface {
	PostActions() []Action
}

func List(ctx context.Context, list sqlx.ScanIterator, pager *pgbuilder.Pager, conditions ...pgbuilder.ConditionBuilder) error {
	d := DBExecutorFromContext(ctx)

	additions := make(builder.Additions, 0)
	where := builder.EmptyCond()

	for i := range conditions {
		condition := conditions[i]

		if ab, ok := condition.(AdditionsBuilder); ok {
			additions = append(additions, ab.ToAdditions(d)...)
		}

		where = builder.And(
			where,
			pgbuilder.ToCondition(d, condition),
		)
	}

	m := list.New()

	if joins, ok := m.(AdditionsBuilder); ok {
		additions = append(additions, joins.ToAdditions(d)...)
	}

	err := pgbuilder.Use(d).
		Select(builder.ColumnsByStruct(m)).
		From(m.(builder.Model)).
		Where(where, additions...).
		List(list, pager)

	if err != nil {
		return err
	}

	if postActions, ok := list.(WithPostActions); ok {
		l := postActions.PostActions()
		for i := range l {
			action := l[i]
			if action != nil {
				if err := action(ctx); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
