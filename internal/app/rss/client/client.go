package client

import (
	"io"
	"net/http"

	"context"

	"github.com/anguloc/zet/internal/dto"
)

type IClient interface {
	Get(ctx context.Context, url string) (*Result, error)
	Post(ctx context.Context, url, contentType string, body io.Reader) (*Result, error)
}

type Result struct {
	HttpResponse *http.Response
}

func New(opts ...OptionFunc) IClient {
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
