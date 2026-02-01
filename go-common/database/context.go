package database

import (
	"context"

	"gorm.io/gorm"
)

const (
	dbContextKey = "__db_instance"
)

func Context(ctx context.Context, db *gorm.DB) context.Context {
	return context.WithValue(ctx, dbContextKey, db)
}

func FromContext(ctx context.Context) *gorm.DB {
	client, ok := ctx.Value(dbContextKey).(*gorm.DB)
	if !ok {
		return nil
	}
	return client
}

func Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return FromContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = Context(ctx, tx)
		return fn(ctx)
	})
}
