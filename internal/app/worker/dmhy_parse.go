package worker

import (
	"context"
	"errors"
	"fmt"

	"github.com/anguloc/zet/internal/app/repo/zetdao/irequest"
	"github.com/anguloc/zet/internal/app/repo/zetdao/irss"
	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/rss/data"
	"github.com/anguloc/zet/internal/app/rss/parse"
	"github.com/anguloc/zet/internal/dto"
	"github.com/anguloc/zet/internal/errx"
	"github.com/anguloc/zet/pkg/log"
	"gorm.io/gorm"
)

type DmhyParse struct {
	requestTable irequest.Repo
	rssTable     irss.Repo
	parse        parse.IParse
}

type dmhyParseCtx struct {
	err       error
	skip      bool
	request   *model.Request
	parseData *data.List
}

func NewDmhyParse() *DmhyParse {
	return &DmhyParse{
		requestTable: irequest.New(),
		rssTable:     irss.New(),
		parse:        parse.New(parse.WithModule(dto.DmhyModule)),
	}
}

func (w *DmhyParse) Run(ctx context.Context) error {
	// 查询request表
	// 锁定处理的数据行
	// 解析
	// 过滤，已存在不再处理
	// 写入rss表
	// 修改request表状态
	dataCtx := &dmhyParseCtx{}
	defer func(e error) {
		w.saveRequest(ctx, dataCtx)
	}(dataCtx.err)

	fns := []func(ctx context.Context, dataCtx *dmhyParseCtx){
		w.queryRequest,
		w.lockRequest,
		w.parseData,
		w.filter,
		w.saveRss,
	}

	for i, fn := range fns {
		fn(ctx, dataCtx)
		if dataCtx.err != nil {
			log.Error(ctx, "DmhyParse_run_err", log.Err(dataCtx.err), log.Int("fn_index", i))
			return dataCtx.err
		} else if dataCtx.skip {
			log.Info(ctx, "DmhyParse_run_skip", log.Int("fn_index", i))
			return nil
		}
	}

	log.Info(ctx, "DmhyParse_run_success", log.Int("save_len", len(dataCtx.parseData.Data)))

	return nil
}

func (w *DmhyParse) queryRequest(ctx context.Context, dataCtx *dmhyParseCtx) {
	dataCtx.request, dataCtx.err = w.requestTable.FirstByMarkStatus(ctx, model.MarkDmHyRss, model.RequestParseInit)
	if errors.Is(dataCtx.err, gorm.ErrRecordNotFound) {
		dataCtx.skip = true
		dataCtx.err = nil
	}
}

func (w *DmhyParse) lockRequest(ctx context.Context, dataCtx *dmhyParseCtx) {
	var count int64
	count, dataCtx.err = w.requestTable.UpdateStatusByIds(ctx, []int64{dataCtx.request.ID}, model.RequestParsing, model.RequestParseInit)
	if dataCtx.err == nil && count != 1 {
		dataCtx.err = errx.ErrLogic.SetMsg(fmt.Sprintf("锁定处理数据行失败,修改结果不为1,count:%d", count))
	}
}

func (w *DmhyParse) parseData(ctx context.Context, dataCtx *dmhyParseCtx) {
	raw := []byte(*dataCtx.request.Content)
	dataCtx.parseData, dataCtx.err = w.parse.SetData(raw).Run(ctx)
}

func (w *DmhyParse) filter(ctx context.Context, dataCtx *dmhyParseCtx) {
	marks := make([]string, 0, len(dataCtx.parseData.Data))
	for _, v := range dataCtx.parseData.Data {
		marks = append(marks, w.rssMark(model.MarkDmHyRss, v.UK))
	}

	rssList, err := w.rssTable.FindByMarkList(ctx, marks)
	if err != nil {
		dataCtx.err = err
		return
	}

	markMap := make(map[string]struct{}, len(dataCtx.parseData.Data))
	for _, rss := range rssList {
		markMap[rss.Mark] = struct{}{}
	}

	newData := make([]*data.Item, 0, len(dataCtx.parseData.Data))
	for _, v := range dataCtx.parseData.Data {
		if _, ok := markMap[w.rssMark(model.MarkDmHyRss, v.UK)]; !ok {
			newData = append(newData, v)
		}
	}
	dataCtx.parseData.Data = newData
}

func (w *DmhyParse) saveRss(ctx context.Context, dataCtx *dmhyParseCtx) {
	if len(dataCtx.parseData.Data) == 0 {
		return
	}
	insertData := make([]*model.Rss, 0, len(dataCtx.parseData.Data))
	for _, v := range dataCtx.parseData.Data {
		insertData = append(insertData, &model.Rss{
			Mark:        w.rssMark(model.MarkDmHyRss, v.UK),
			Title:       v.Title,
			Link:        v.Link,
			Author:      v.Author,
			PubDate:     &v.PubDate,
			Desctiption: &v.Description,
			Bittorrent:  &v.Bittorrent,
			GUID:        v.Guid,
			Uk:          v.UK,
		})
	}
	var count int64
	count, dataCtx.err = w.rssTable.BatchInsertRss(ctx, insertData)
	if dataCtx.err != nil {
		return
	}

	if int(count) != len(dataCtx.parseData.Data) {
		dataCtx.err = errx.ErrLogic.SetMsg(fmt.Sprintf("实际插入数量与应该插入数量不一致,,count:%d,data_len:%d", count, len(dataCtx.parseData.Data)))
		return
	}
}

func (w *DmhyParse) saveRequest(ctx context.Context, dataCtx *dmhyParseCtx) {
	// 没到查询错误这一步就不流转
	if dataCtx.request == nil || dataCtx.request.ID <= 0 {
		return
	}
	newStatus := model.RequestParseSuccess
	if dataCtx.err != nil {
		newStatus = model.RequestParseFail
	}
	count, err := w.requestTable.UpdateStatusByIds(ctx, []int64{dataCtx.request.ID}, int32(newStatus), model.RequestParsing)
	if err != nil || count != 1 {
		log.Error(ctx, "DmhyParse_saveRequest_err", log.Err(err), log.Int64("count", count))
	}
}

func (w *DmhyParse) rssMark(mark, id string) string {
	return fmt.Sprintf("%s_%s", mark, id)
}
