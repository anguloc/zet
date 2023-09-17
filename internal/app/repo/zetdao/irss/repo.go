package irss

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
)

//go:generate mockgen -source $GOFILE -destination rss_mock.go -package $GOPACKAGE

type Repo interface {
	Reader
	Writer
}

type Reader interface {
	FirstByMark(ctx context.Context, mark string) (*model.Rss, error)
}

type Writer interface {
}

type Rss struct {
}

var (
	_ Repo   = &Rss{}
	_ Reader = &Rss{}
	_ Writer = &Rss{}
)

func New() *Rss {
	return &Rss{}
}
