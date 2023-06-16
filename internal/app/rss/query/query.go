package query

import "context"

type IQuery interface {
	Get(ctx context.Context) (*Result, error)
}

type Result struct {
}
