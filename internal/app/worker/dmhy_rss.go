package worker

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/anguloc/zet/internal/app/repo/kv"
	"github.com/anguloc/zet/internal/pkg/log"
	"github.com/robfig/cron/v3"
)

type DmhyRss struct {
	cron *cron.Cron
}

func (w *DmhyRss) Init(ctx context.Context) error {
	w.cron = cron.New(cron.WithLogger(cron.DiscardLogger))
	return nil
}

func (w *DmhyRss) Run(ctx context.Context) error {
	ch := make(chan struct{})
	_, err := w.cron.AddFunc("* * * * *", func() {
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
	<-ch
	return nil
}

func (w *DmhyRss) handle(ctx context.Context) {
	repo := kv.New()
	// 查询rss svr
	if err := w.queryResource(ctx, repo, w.getRssUrlList()); err != nil {
		log.Error(ctx, "queryRssErr", log.NamedError("err", err))
		return
	}

	// 解析资源
	//data, err := w.parseResource(ctx, repo, w.getRssUrlList())
	//if err != nil {
	//
	//}

	// 获取所有关键词
	//keywords := []string{"我推的孩子"}

	// 匹配资源
	// 异步获取
}

func (w *DmhyRss) getRssUrlList() map[string]string {
	return map[string]string{
		"dmhy": "https://www.dmhy.org/topics/rss/rss.xml",
	}
}

func (w *DmhyRss) queryRss(ctx context.Context, url string) ([]byte, int, error) {
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				hosts := map[string]string{
					"www.dmhy.org:443": "172.67.98.15:443",
				}
				if n, ok := hosts[addr]; ok {
					addr = n
				}
				d := &net.Dialer{
					Timeout: time.Second * 10,
				}
				return d.DialContext(ctx, network, addr)
			},
		},
	}
	rsp, err := client.Get(url)
	if err != nil {
		return nil, rsp.StatusCode, err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)

	return body, rsp.StatusCode, err
}

func (w *DmhyRss) queryResource(ctx context.Context, repo kv.IKV, rss map[string]string) error {
	oneSuccess := false
	for name, url := range rss {
		body, code, err := w.queryRss(ctx, url)
		if err != nil {
			log.Error(ctx, "httpRequestErr", log.NamedError("err", err))
			continue
		}
		if code != 200 {
			log.Error(ctx, "quertRssCodeNot200", log.Int("code", code))
			continue
		}

		if err := repo.Put(name, body); err != nil {
			log.Error(ctx, "rssSaveErr", log.NamedError("err", err))
			continue
		}
		oneSuccess = true
	}

	if !oneSuccess {
		return fmt.Errorf("queryRssNotSucc")
	}

	return nil
}

func (w *DmhyRss) parseResource(ctx context.Context, repo kv.IKV, rss map[string]string) (map[string]*rssData, error) {
	res := make(map[string]*rssData, len(rss))
	for name := range rss {
		value, err := repo.Get(name)
		if err != nil {
			continue
		}
		data := &rssData{}
		if err := xml.Unmarshal(value, data); err != nil {
			log.Error(ctx, "xmlParseErr", log.NamedError("err", err))
			continue
		}

	}

	return res, nil
}
