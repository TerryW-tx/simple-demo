package dal

import (
	"context"
)

type ctxTransaction struct{}

func GetQueryByCtx(ctx context.Context) *Query {
	dbI := ctx.Value(ctxTransaction{})

	if dbI != nil {
		db, ok := dbI.(*Query)
		if !ok {
			panic("unexpected context query value type")
		}
		if db != nil {
			return db
		}
	}
	return Q
}

func SetCtxQuery(ctx context.Context, q *Query) context.Context {
	return context.WithValue(ctx, ctxTransaction{}, q)
}

func Transaction(ctx context.Context, fn func(txCtx context.Context) error) error {
	q := GetQueryByCtx(ctx)
	return q.Transaction(func(tx *Query) error {
		txCtx := SetCtxQuery(ctx, tx)
		return fn(txCtx)
	})
}
