package client

import (
	"context"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/anguloc/zet/pkg/conf"
)

type Dmhy struct {
}

func NewDmhy() *Dmhy {
	return &Dmhy{}
}

func (d *Dmhy) Get(ctx context.Context, url string) (*Result, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	rsp, err := d.client().Do(req)
	if err != nil {
		return nil, err
	}
	return &Result{HttpResponse: rsp}, nil
}

func (d *Dmhy) Post(ctx context.Context, url, contentType string, body io.Reader) (*Result, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	rsp, err := d.client().Do(req)
	if err != nil {
		return nil, err
	}
	return &Result{HttpResponse: rsp}, nil
}

func (d *Dmhy) client() *http.Client {
	return &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				hosts := conf.Conf().Dmhy.Host
				if n, ok := hosts[addr]; ok {
					addr = n
				}
				di := &net.Dialer{
					Timeout: time.Second * 10,
				}
				return di.DialContext(ctx, network, addr)
			},
		},
	}
}
