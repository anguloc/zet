package worker

import (
	"testing"

	"github.com/anguloc/zet/internal/app/repo/zetdao/irequest"
	"github.com/anguloc/zet/internal/app/repo/zetdao/irss"
	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/rss/data"
	"github.com/anguloc/zet/internal/app/rss/parse"
	"github.com/anguloc/zet/pkg/application"
	"github.com/anguloc/zet/pkg/testx"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	if err := application.New().Init(testx.Context(nil), application.WithOptionConfig("")); err != nil {
		panic(err)
	}
	m.Run()
}

// go test ./internal/app/worker/ -v -run TestDmhyParse_run$ --count=1
func TestDmhyParse_run(t *testing.T) {
	ctx := testx.Context(nil)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	reqTable := irequest.NewMockRepo(ctrl)
	reqTable.EXPECT().
		FirstByMarkStatus(gomock.Any(), model.MarkDmHyRss, int32(model.RequestParseInit)).
		Return(&model.Request{
			ID:          1,
			Mark:        model.MarkDmHyRss,
			ParseStatus: model.RequestParseInit,
			Content:     new(string),
		}, nil).
		Times(1)
	reqTable.EXPECT().
		UpdateStatusByIds(gomock.Any(), []int64{1}, int32(model.RequestParsing), int32(model.RequestParseInit)).
		Return(int64(1), nil).
		Times(1)
	reqTable.EXPECT().
		UpdateStatusByIds(gomock.Any(), []int64{1}, int32(model.RequestParseSuccess), int32(model.RequestParsing)).
		Return(int64(1), nil).
		Times(1)

	rssTable := irss.NewMockRepo(ctrl)
	rssTable.EXPECT().
		FindByMarkList(gomock.Any(), []string{"dmhy_rss_test_a_uk", "dmhy_rss_test_b_uk"}).
		Return([]*model.Rss{
			{
				ID:   1,
				Mark: "dmhy_rss_test_a_uk",
			},
		}, nil).
		Times(1)
	rssTable.EXPECT().
		BatchInsertRss(gomock.Any(), gomock.Any()).
		Return(int64(1), nil).
		Times(1)

	parser := parse.NewMockIParse(ctrl)
	parser.EXPECT().SetData(gomock.Any()).Return(parser).Times(1)
	parser.EXPECT().Run(gomock.Any()).Return(&data.List{
		Data: []*data.Item{
			{
				Title:  "test_a",
				Link:   "test_link",
				Author: "test_author",
				UK:     "test_a_uk",
			},
			{
				Title:  "test_b",
				Link:   "test_link",
				Author: "test_author",
				UK:     "test_b_uk",
			},
		},
	}, nil).Times(1)

	d := &DmhyParse{
		requestTable: reqTable,
		rssTable:     rssTable,
		parse:        parser,
	}
	//d = newDmhyParse()
	err := d.Run(ctx)
	assert.Nil(t, err)
}
