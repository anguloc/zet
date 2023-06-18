package parse

import (
	"context"
	"fmt"
)

type Default struct{}

var ModuleMissingErr = fmt.Errorf("failed to missing module")

func NewDefault() *Default {
	return &Default{}
}

func (d *Default) SetData(bytes []byte) error {
	return ModuleMissingErr
}

func (d *Default) Run(ctx context.Context) (*Result, error) {
	return nil, ModuleMissingErr
}
