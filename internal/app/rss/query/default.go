package query

import (
	"context"
	"fmt"
)

type Default struct{}

func NewDefault() *Default {
	return &Default{}
}

var ModuleMissingErr = fmt.Errorf("failed to missing module")

func (d *Default) Command(ctx context.Context) error {
	return ModuleMissingErr
}

func (d *Default) Get(ctx context.Context) (*Result, error) {
	return nil, ModuleMissingErr
}
