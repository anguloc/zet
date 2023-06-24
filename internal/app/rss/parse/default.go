package parse

import (
	"context"
	"fmt"

	"github.com/anguloc/zet/internal/app/rss"
)

type Default struct{}

var ModuleMissingErr = fmt.Errorf("failed to missing module")

func NewDefault() *Default {
	return &Default{}
}

func (d *Default) SetData(bytes []byte) IParse {
	return d
}

func (d *Default) Run(ctx context.Context) (*rss.List, error) {
	return nil, ModuleMissingErr
}
