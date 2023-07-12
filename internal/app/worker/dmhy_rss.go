package worker

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/anguloc/zet/internal/app/rss/client"
	"github.com/anguloc/zet/internal/dao"
	"github.com/anguloc/zet/internal/dao/zet_model"
	"github.com/anguloc/zet/internal/dao/zet_query"
	"github.com/anguloc/zet/internal/dto"
	"github.com/anguloc/zet/pkg/cron"
	"github.com/anguloc/zet/pkg/log"
)

type DmhyRss struct {
	client       client.IClient
	requestTable zet_query.IRequestDo
}

func Init() {
	_ = cron.NewCron().RegisterJob("3 * * * *", NewDmhyRss())
}

func NewDmhyRss() *DmhyRss {
	return &DmhyRss{
		client:       client.New(client.WithModule(dto.DmhyModule)),
		requestTable: dao.Zet().Request.WithContext(nil),
	}
}

func (w *DmhyRss) Run(ctx context.Context) error {
	url := "https://www.dmhy.org/topics/rss/rss.xml"

	rsp, err := w.client.Get(ctx, url)
	if err != nil {
		log.Error(ctx, "dmhyRssReqErr", log.NamedError("err", err))
		return err
	}
	defer func() { _ = rsp.HttpResponse.Body.Close() }()
	if rsp.HttpResponse.StatusCode != http.StatusOK {
		log.Error(ctx, "dmhyRssReqCodeErr", log.Int("http_code", rsp.HttpResponse.StatusCode))
		return fmt.Errorf("rss http err")
	}
	body, err := io.ReadAll(rsp.HttpResponse.Body)
	if err != nil {
		log.Error(ctx, "dmhyRssReadBodyErr", log.NamedError("err", err))
		return err
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
		return err
	}
	log.Info(ctx, "dmhyRssRequestSuccess")
	return nil
}
