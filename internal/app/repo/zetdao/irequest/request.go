package irequest

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/query"
	"github.com/anguloc/zet/pkg/db"
)

func (r *Request) Insert(ctx context.Context, data *model.Request) error {
	conn, _ := db.Zet()
	q := query.Use(conn)
	return q.Request.WithContext(ctx).WriteDB().Create(data)
}

func (r *Request) FirstByMarkStatus(ctx context.Context, mark string, status int32) (*model.Request, error) {
	q := query.Request
	return q.WithContext(ctx).ReadDB().
		Where(q.Mark.Eq(mark)).
		Where(q.ParseStatus.Eq(status)).
		First()
}
