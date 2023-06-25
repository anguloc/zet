package parse

import (
	"context"
	"fmt"

	"github.com/anguloc/zet/internal/app/rss/data"
)

type Default struct{}

var ModuleMissingErr = fmt.Errorf("failed to missing module")

func NewDefault() *Default {
	return &Default{}
}

func (d *Default) SetData(bytes []byte) IParse {
	return d
}

func (d *Default) Run(ctx context.Context) (*data.List, error) {
	return nil, ModuleMissingErr
}
