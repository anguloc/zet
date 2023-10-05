package irss

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/query"
	"gorm.io/gen"
)

func (r *Rss) BatchInsertRss(ctx context.Context, values []*model.Rss) (int64, error) {
	q := query.Rss
	res := q.WithContext(ctx).WriteDB().
		WithResult(func(tx gen.Dao) {
			_ = tx.Create(values)
		})
	return res.RowsAffected, res.Error
}
