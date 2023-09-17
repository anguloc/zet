package worker

import (
	"context"
	"errors"
	"fmt"

	"github.com/anguloc/zet/internal/app/repo/zetdao/irequest"
	"github.com/anguloc/zet/internal/app/repo/zetdao/irss"
	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/pkg/log"
	"gorm.io/gorm"
)

type DmhyParse struct {
	requestTable irequest.Reader
	rssTable     irss.Repo
}

type dmhyParseCtx struct {
	err     error
	skip    bool
	request *model.Request
}

func newDmhyParse() *DmhyParse {
	return &DmhyParse{
		requestTable: irequest.New(),
		//rssTable:     irss.New(),
	}
}

func (w *DmhyParse) Run(ctx context.Context) error {
	_ = w.run(ctx, 1)
	return nil
}

func (w *DmhyParse) run(ctx context.Context, times int) error {
	if times >= 300 {
		log.Info(ctx, "DmhyParse_FindByMarkStatus_empty", log.Int("times", times))
		return nil
	}
	// 查询request表
	// 解析
	// 过滤，已存在不再处理
	// 写入rss表
	// 修改request表状态
	fns := []func(ctx context.Context, data *dmhyParseCtx){
		w.queryRequest,
		w.parse,
		w.filter,
		w.saveRss,
		w.saveRequest,
	}

	data := &dmhyParseCtx{}
	for i, fn := range fns {
		fn(ctx, data)
		if data.err != nil {
			log.Error(ctx, "DmhyParse_FindByMarkStatus_err", log.Err(data.err), log.Int("fn_index", i), log.Int("times", times))
			return data.err
		} else if data.skip {
			log.Info(ctx, "DmhyParse_FindByMarkStatus_empty", log.Int("fn_index", i), log.Int("times", times))
			return nil
		}
	}

	times++
	return w.run(ctx, times)
}

func (w *DmhyParse) queryRequest(ctx context.Context, data *dmhyParseCtx) {
	data.request, data.err = w.requestTable.FirstByMarkStatus(ctx, model.MarkDmHyRss, model.RequestParseInit)
	if errors.Is(data.err, gorm.ErrRecordNotFound) {
		data.skip = true
		data.err = nil
	}
}

func (w *DmhyParse) parse(ctx context.Context, data *dmhyParseCtx) {
}

func (w *DmhyParse) filter(ctx context.Context, data *dmhyParseCtx) {
}

func (w *DmhyParse) saveRss(ctx context.Context, data *dmhyParseCtx) {
}

func (w *DmhyParse) saveRequest(ctx context.Context, data *dmhyParseCtx) {
}

func (w *DmhyParse) rssMark(mark string, id int64) string {
	return fmt.Sprintf("%s_%d", mark, id)
}
