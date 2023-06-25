package rss

import (
	"context"
	"fmt"
	"time"

	"github.com/anguloc/zet/internal/app/rss/data"
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

var MissDataErr = fmt.Errorf("failed to miss data")

func (d *Dmhy) Run(ctx context.Context) error {
	queryRes, err := d.query.Get(ctx)
	if err != nil {
		return err
	}

	parseRes, err := d.parse.SetData(queryRes.Data).Run(ctx)
	if err != nil {
		return err
	}

	list, err := d.filter(parseRes)
	if err != nil {
		return err
	}

	if err = d.push(list); err != nil {
		return err
	}

	return nil
}

func (d *Dmhy) filter(list *data.List) (*data.TransmissionList, error) {
	var err error
	if len(list.Data) == 0 {
		return nil, MissDataErr
	}
	res := &data.TransmissionList{
		StartTime: time.Now().Unix(),
	}

	fns := []func(*data.Item) error{
		d.filterTitle,
		d.filterAuthor,
	}

	for _, v := range list.Data {
		for _, fn := range fns {
			if err = fn(v); err != nil {
				continue
			}
		}
	}

	if len(res.Data) == 0 {
		return nil, MissDataErr
	}

	return nil, nil
}

func (d *Dmhy) filterTitle(item *data.Item) error {
	return nil
}

func (d *Dmhy) filterAuthor(item *data.Item) error {
	return nil
}

func (d *Dmhy) push(list *data.TransmissionList) error {
	return nil
}
