package parse

import (
	"context"
	"github.com/anguloc/zet/internal/dto"
)

type IParse interface {
	SetData(data []byte) error
	Run(ctx context.Context) (*Result, error)
}

type Result struct {
}

func New(opts ...OptionFunc) IParse {
	opt := &Option{}
	for _, fn := range opts {
		fn(opt)
	}
	switch opt.module {
	case dto.DmhyModule:
		return NewDmhy()
	}

	return NewDefault()
}
