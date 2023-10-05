package parse

import (
	"context"

	"github.com/anguloc/zet/internal/app/rss/data"
	"github.com/anguloc/zet/internal/dto"
)

//go:generate mockgen -source $GOFILE -destination parse_mock.go -package $GOPACKAGE

type IParse interface {
	SetData(data []byte) IParse
	Run(ctx context.Context) (*data.List, error)
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
