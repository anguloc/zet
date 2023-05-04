package application

import (
	"context"
)

type IWorker interface {
	Init(ctx context.Context) error
	Run(ctx context.Context) error
}
