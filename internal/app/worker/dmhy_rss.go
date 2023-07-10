package worker

import (
	"context"
	"io"
	"net/http"

	"github.com/anguloc/zet/internal/app/rss/client"
	"github.com/anguloc/zet/internal/dao"
	"github.com/anguloc/zet/internal/dao/zet_model"
	"github.com/anguloc/zet/internal/dao/zet_query"
	"github.com/anguloc/zet/internal/dto"
	"github.com/anguloc/zet/internal/pkg/log"
	"github.com/robfig/cron/v3"
)

type DmhyRss struct {
	cron         *cron.Cron
	client       client.IClient
	requestTable zet_query.IRequestDo
}

func (w *DmhyRss) Init(ctx context.Context) error {
	w.cron = cron.New(cron.WithLogger(cron.DiscardLogger))
	w.client = client.New(client.WithModule(dto.DmhyModule))
	w.requestTable = dao.Zet().Request.WithContext(ctx)
	return nil
}

func (w *DmhyRss) Run(ctx context.Context) error {
	_, err := w.cron.AddFunc("3 * * * *", func() {
		w.handle(ctx)
	})
	if err != nil {
		return err
	}
	w.cron.Start()
	select {
	case <-ctx.Done():
		w.cron.Stop()
	}
	return nil
}

func (w *DmhyRss) handle(ctx context.Context) {
	url := "https://www.dmhy.org/topics/rss/rss.xml"

	rsp, err := w.client.Get(ctx, url)
	if err != nil {
		log.Error(ctx, "dmhyRssReqErr", log.NamedError("err", err))
		return
	}
	defer func() { _ = rsp.HttpResponse.Body.Close() }()
	if rsp.HttpResponse.StatusCode != http.StatusOK {
		log.Error(ctx, "dmhyRssReqCodeErr", log.Int("http_code", rsp.HttpResponse.StatusCode))
		return
	}
	body, err := io.ReadAll(rsp.HttpResponse.Body)
	if err != nil {
		log.Error(ctx, "dmhyRssReadBodyErr", log.NamedError("err", err))
		return
	}
	str := string(body)
	err = w.requestTable.WithContext(ctx).Create(&zet_model.Request{
		URL:        url,
		Mark:       "dmhy_rss",
		RequestNum: 1,
		Content:    &str,
	})
	if err != nil {
		log.Error(ctx, "dmhyRssInsertDbErr", log.NamedError("err", err))
		return
	}
	log.Info(ctx, "dmhyRssRequestSuccess")
}
