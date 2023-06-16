package query

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"
)

var dmhyUrl = []string{
	"https://www.dmhy.org/topics/rss/rss.xml",
}

type Dmhy struct {
	url string
}

func NewDmhy() IQuery {
	return &Dmhy{}
}

func (q *Dmhy) Get(ctx context.Context) (*Result, error) {
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
	rsp, err := client.Get(q.getUrl())
	if err != nil {
		return nil, rsp.StatusCode, err
	}
	defer rsp.Body.Close()

	body, err := io.ReadAll(rsp.Body)

	return body, rsp.StatusCode, err
}

func (q *Dmhy) getUrl() string {
	if q.url == "" {
		return dmhyUrl[0]
	}
	return q.url
}
