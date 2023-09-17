package request

import (
	"context"

	"github.com/anguloc/zet/internal/app/repo/zetdao/model"
	"github.com/anguloc/zet/internal/app/repo/zetdao/query"
	"github.com/anguloc/zet/pkg/db"
)

type Request struct {
}

func (r *Request) InsertRequest(ctx context.Context, data *model.Request) error {
	conn, _ := db.Zet()
	q := query.Use(conn)
	return q.Request.WithContext(ctx).Create(data)
}
