package irss

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/query"
)

func (r *Rss) FindByMarkList(ctx context.Context, mark []string) ([]*model.Rss, error) {
	q := query.Rss
	return q.WithContext(ctx).ReadDB().
		Where(q.Mark.In(mark...)).
		Find()
}
