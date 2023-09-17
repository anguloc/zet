package irequest

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
)

//go:generate mockgen -source $GOFILE -destination request_mock.go -package $GOPACKAGE

type Repo interface {
	Reader
	Writer
}

type Reader interface {
	FirstByMarkStatus(ctx context.Context, mark string, status int32) (*model.Request, error)
}

type Writer interface {
	Insert(ctx context.Context, data *model.Request) error
}

type Request struct {
}

var (
	_ Repo   = &Request{}
	_ Reader = &Request{}
	_ Writer = &Request{}
)

func New() *Request {
	return &Request{}
}
