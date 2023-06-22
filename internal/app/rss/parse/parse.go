package parse

import (
	"context"

	"github.com/anguloc/zet/internal/dto"
	"github.com/anguloc/zet/internal/dto/rss"
)

type IParse interface {
	SetData(data []byte) IParse
	Run(ctx context.Context) (*rss.List, error)
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
