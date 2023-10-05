package irequest

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/query"
)

func (r *Request) Insert(ctx context.Context, data *model.Request) error {
	return query.Q.Request.WithContext(ctx).WriteDB().Create(data)
}

func (r *Request) UpdateStatusByIds(ctx context.Context, ids []int64, newStatus, oldStatus int32) (int64, error) {
	q := query.Request
	res, err := q.WithContext(ctx).WriteDB().
		Where(q.ID.In(ids...)).
		Where(q.ParseStatus.Eq(oldStatus)).
		UpdateColumn(q.ParseStatus, newStatus)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}
