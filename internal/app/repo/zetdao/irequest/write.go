package irequest

import (
	"context"
	"time"

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

func (r *Request) CleanOldData(ctx context.Context, mark string, status int32, limitTime time.Time) (int64, error) {
	q := query.Request
	res, err := q.WithContext(ctx).WriteDB().
		Where(q.Mark.Eq(mark)).
		Where(q.ParseStatus.Eq(status)).
		Where(q.CreatedAt.Lt(limitTime)).
		Delete()
	if err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}
