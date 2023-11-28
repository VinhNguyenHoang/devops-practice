package database

import "context"

type ConnManager interface {
	Init(ctx context.Context) error
	Query(ctx context.Context, args ...interface{}) error
}
