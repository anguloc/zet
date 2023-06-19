package rss

import (
	"context"

	"github.com/anguloc/zet/internal/app/rss/parse"
	"github.com/anguloc/zet/internal/app/rss/query"
	"github.com/anguloc/zet/internal/dto"
)

type Dmhy struct {
	query query.IQuery
	parse parse.IParse
}

func NewDmhy() *Dmhy {
	return &Dmhy{
		query: query.New(query.WithModule(dto.DmhyModule)),
		parse: parse.New(parse.WithModule(dto.DmhyModule)),
	}
}

func (d *Dmhy) Run(ctx context.Context) error {
	queryRes, err := d.query.Get(ctx)
	if err != nil {
		return err
	}
	parseRes, err := d.parse.SetData(queryRes.Data).Run(ctx)
	if err != nil {
		return err
	}

	return nil
}
