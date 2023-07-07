package rss

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/anguloc/zet/internal/app/rss/parse"
	"github.com/anguloc/zet/internal/app/rss/query"
	"github.com/anguloc/zet/internal/dto"
	"github.com/anguloc/zet/internal/pkg/console"
	"github.com/anguloc/zet/internal/pkg/safe"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "rss",
	Short: "rss",
	Long:  "Really Simple Syndication",
	Run:   Run,
}

func init() {
}

func Run(cmd *cobra.Command, args []string) {
	console.Info("rss start")
	pCtx := cmd.Context()
	ctx, cancel := context.WithCancel(pCtx)
	defer cancel()

	biz := struct {
		query query.IQuery
		parse parse.IParse
	}{
		query: query.New(query.WithModule(dto.DmhyModule)),
		parse: parse.New(parse.WithModule(dto.DmhyModule)),
	}

	fileData, _ := os.ReadFile(safe.Path("internal/app/rss/parse/testdata/dmhy.rss"))
	ctrl := gomock.NewController(&testing.T{})
	q := query.NewMockIQuery(ctrl)
	biz.query = q
	q.EXPECT().Get(gomock.Any()).AnyTimes().Return(&query.Result{
		HttpStatusCode: 200,
		Data:           fileData,
	}, nil)

	console.Info("query start")
	data, err := biz.query.Get(ctx)
	if err != nil {
		console.Errorf("failed to query,err info:[%s]\n", err)
		return
	}
	console.Info("query end")

	console.Info("parse start")
	res, err := biz.parse.SetData(data.Data).Run(ctx)
	if err != nil {
		console.Errorf("failed to parse,err info:[%s]\n", err)
		return
	}
	console.Info("parse end")

	pattern := "我推"
	for _, v := range res.Data {
		if strings.Contains(v.Title, pattern) {
			continue
		}
	}

	console.Info("rss end")
}
