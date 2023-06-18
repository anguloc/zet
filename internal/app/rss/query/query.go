package query

//go:generate mockgen -source ./$GOFILE -destination ./query_mock.go -package=$GOPACKAGE

import (
	"context"
	"github.com/anguloc/zet/internal/dto"
)

type IQuery interface {
	Command(ctx context.Context) error
	Get(ctx context.Context) (*Result, error)
}

type Result struct {
	HttpStatusCode int
	Data           []byte
}

func New(opts ...OptionFunc) IQuery {
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
