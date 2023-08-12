package worker

import (
	"context"

	"github.com/anguloc/zet/internal/dao/zet_query"
)

type DmhyParse struct {
	requestTable zet_query.IRequestDo
	rssTable     zet_query.IRssDo
}

func newDmhyParse() *DmhyParse {
	return &DmhyParse{
		requestTable: nil,
		rssTable:     nil,
	}
}

func (w *DmhyParse) Run(ctx context.Context) error {
	return nil
}
