package query

import (
	"context"
	"io"

	"github.com/anguloc/zet/internal/app/rss/client"
	"github.com/anguloc/zet/internal/dto"
)

var dmhyUrl = []string{
	"https://www.dmhy.org/topics/rss/rss.xml",
}

type Dmhy struct {
	url    string
	client client.IClient
}

func NewDmhy() *Dmhy {
	return &Dmhy{
		client: client.New(client.WithModule(dto.DmhyModule)),
	}
}

func (d *Dmhy) Command(ctx context.Context, flag string) error {
	d.url = flag
	return nil
}

func (d *Dmhy) Get(ctx context.Context) (*Result, error) {
	res, err := d.client.Get(ctx, d.getUrl())
	if err != nil {
		return nil, err
	}
	rsp := res.HttpResponse
	defer func() { _ = rsp.Body.Close() }()

	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, err
	}
	return &Result{
		HttpStatusCode: rsp.StatusCode,
		Data:           body,
	}, nil
}

func (d *Dmhy) getUrl() string {
	if d.url == "" {
		return dmhyUrl[0]
	}
	return d.url
}
