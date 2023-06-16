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

type AnimationRss struct {
	isRunning bool
	cron      *cron.Cron
}

func (w *AnimationRss) Init(ctx context.Context) error {
	fmt.Println("init animation rss")
	return nil
}

func (w *AnimationRss) Run(ctx context.Context) error {
	log.Info(ctx, "AnimationRss_Start")
	ch := make(chan struct{})
	w.isRunning = true
	go func() {
		for w.isRunning {

		}
		close(ch)
	}()

	select {
	case <-ctx.Done():
		w.isRunning = false
	}
	<-ch

	log.Info(ctx, "AnimationRss_End")

	return nil
}

func (w *AnimationRss) handle(ctx context.Context) {
	repo := kv.New()
	// 查询rss svr
	if err := w.queryResource(ctx, repo, w.getRssUrlList()); err != nil {
		log.Error(ctx, "queryRssErr", log.NamedError("err", err))
		return
	}

	// 解析资源
	data, err := w.parseResource(ctx, repo, w.getRssUrlList())
	if err != nil {

	}

	// 获取所有关键词
	keywords := []string{"我推的孩子"}

	// 匹配资源
	// 异步获取
}

func (w *AnimationRss) getRssUrlList() map[string]string {
	return map[string]string{
		"dmhy": "https://www.dmhy.org/topics/rss/rss.xml",
	}
}

func (w *AnimationRss) queryRss(ctx context.Context, url string) ([]byte, int, error) {
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

func (w *AnimationRss) queryResource(ctx context.Context, repo kv.IKV, rss map[string]string) error {
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

func (w *AnimationRss) parseResource(ctx context.Context, repo kv.IKV, rss map[string]string) (map[string]*rssData, error) {
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

type rssData struct {
	XMLName xml.Name `xml:"rss"`
	Text    string   `xml:",chardata"`
	Version string   `xml:"version,attr"`
	Content string   `xml:"content,attr"`
	Wfw     string   `xml:"wfw,attr"`
	Channel struct {
		Text        string `xml:",chardata"`
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		PubDate     string `xml:"pubDate"`
		Item        []struct {
			Text        string `xml:",chardata"`
			Title       string `xml:"title"`
			Link        string `xml:"link"`
			PubDate     string `xml:"pubDate"`
			Description string `xml:"description"`
			Enclosure   struct {
				Text   string `xml:",chardata"`
				URL    string `xml:"url,attr"`
				Length string `xml:"length,attr"`
				Type   string `xml:"type,attr"`
			} `xml:"enclosure"`
			Author string `xml:"author"`
			Guid   struct {
				Text        string `xml:",chardata"`
				IsPermaLink string `xml:"isPermaLink,attr"`
			} `xml:"guid"`
			Category struct {
				Text   string `xml:",chardata"`
				Domain string `xml:"domain,attr"`
			} `xml:"category"`
		} `xml:"item"`
	} `xml:"channel"`
}
