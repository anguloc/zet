package dao

import (
	"github.com/anguloc/zet/internal/dao/zet_query"
	"github.com/anguloc/zet/internal/pkg/db"
)

func Zet() *zet_query.Query {
	db := db.Zet()
	return zet_query.Use(db)
}
