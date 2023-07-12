package dao

import (
	"github.com/anguloc/zet/internal/dao/zet_query"
	"github.com/anguloc/zet/pkg/db"
)

func Zet() *zet_query.Query {
	conn := db.Zet()
	return zet_query.Use(conn)
}
