package zet

import (
	"context"

	"github.com/anguloc/zet/internal/dao/zet_model"
)

//go:generate mockgen -destination gen_mock.go -package zet_query . IRssDo

type Reader interface {
	CreateRequestRow(ctx context.Context, request zet_model.Request)
}

type Writer interface {
}

type Option struct {
}
