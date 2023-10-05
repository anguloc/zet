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
	// FirstByMarkStatus 根据标识以及状态查询一条数据
	FirstByMarkStatus(ctx context.Context, mark string, status int32) (*model.Request, error)
}

type Writer interface {
	// Insert 写一条数据
	Insert(ctx context.Context, data *model.Request) error

	// UpdateStatusByIds 根据id和状态更新一条数据
	UpdateStatusByIds(ctx context.Context, ids []int64, newStatus, oldStatus int32) (int64, error)
}

var (
	_ Repo   = &Request{}
	_ Reader = &Request{}
	_ Writer = &Request{}
)

type Request struct {
}

func New() *Request {
	return &Request{}
}
