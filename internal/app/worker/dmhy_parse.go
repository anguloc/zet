package worker

import (
	"context"

	"github.com/anguloc/zet/internal/dao"
	"github.com/anguloc/zet/internal/dao/zet_query"
)

type DmhyParse struct {
	requestTable zet_query.IRequestDo
	rssTable     zet_query.IRssDo
}

func newDmhyParse() *DmhyParse {
	return &DmhyParse{
		requestTable: dao.Zet().Request.WithContext(nil),
		rssTable:     dao.Zet().Rss.WithContext(nil),
	}
}

func (w *DmhyParse) Run(ctx context.Context) error {
	return nil
}
