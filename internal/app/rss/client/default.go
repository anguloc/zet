package client

import (
	"context"
	"io"
	"net/http"
)

type Default struct{}

func NewDefault() *Default {
	return &Default{}
}

func (d *Default) Get(ctx context.Context, url string) (*Result, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return &Result{HttpResponse: rsp}, nil
}

func (d *Default) Post(ctx context.Context, url, contentType string, body io.Reader) (*Result, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	client := &http.Client{}
	rsp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return &Result{HttpResponse: rsp}, nil
}
