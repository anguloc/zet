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
	"github.com/anguloc/zet/internal/pkg/cron"
	"github.com/anguloc/zet/internal/pkg/log"
)

type DmhyRss struct {
	cron         cron.ICron
	client       client.IClient
	requestTable zet_query.IRequestDo
}

func (w *DmhyRss) Init(ctx context.Context) error {
	w.cron = cron.NewCron()
	w.client = client.New(client.WithModule(dto.DmhyModule))
	w.requestTable = dao.Zet().Request.WithContext(ctx)
	return w.cron.RegisterFunc(ctx, "3 * * * *", func(ctx context.Context) error {
		w.handle(ctx)
		return nil
	})
}

func (w *DmhyRss) Run(ctx context.Context) error {
	return w.cron.Start(ctx)
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
