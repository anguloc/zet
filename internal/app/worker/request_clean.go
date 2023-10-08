package worker

import (
	"context"
	"time"

	"github.com/anguloc/zet/internal/app/repo/zetdao/irequest"
	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/pkg/log"
)

type RequestClean struct {
	requestTable irequest.Repo
}

func NewRequestClean() *RequestClean {
	return &RequestClean{
		requestTable: irequest.New(),
	}
}

func (r *RequestClean) Run(ctx context.Context) error {
	limitTime := time.Now().AddDate(0, -1, 0)
	count, err := r.requestTable.CleanOldData(ctx, model.MarkDmHyRss, model.RequestParseSuccess, limitTime)
	if err != nil {
		log.Error(ctx, "RequestClean_err", log.Err(err), log.Time("limitTime", limitTime))
		return nil
	}
	log.Info(ctx, "RequestClean_success", log.Time("limitTime", limitTime), log.Int64("count", count))
	return nil
}
