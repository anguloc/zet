package irequest

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/query"
)

func (r *Request) FirstByMarkStatus(ctx context.Context, mark string, status int32) (*model.Request, error) {
	q := query.Request
	return q.WithContext(ctx).ReadDB().
		Where(q.Mark.Eq(mark)).
		Where(q.ParseStatus.Eq(status)).
		First()
}
