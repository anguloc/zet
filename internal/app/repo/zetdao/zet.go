package zetdao

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/request"
)

//go:generate mockgen -destination gen_mock.go -package zet_query . IRssDo

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
