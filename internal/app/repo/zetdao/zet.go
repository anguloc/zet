package zetdao

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/request"
)

//go:generate mockgen -source $GOFILE -destination zet_mock.go -package $GOPACKAGE

type Repo interface {
	Reader
	Writer
}

type Reader interface {
}

type Writer interface {
	InsertRequest(ctx context.Context, data *model.Request) error
}

type Impl struct {
	request.Request
}

func New() *Impl {
	return &Impl{}
}
