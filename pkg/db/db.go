package db

import (
	"context"
	"sync"
	"time"

	"github.com/anguloc/zet/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbMap = make(map[string]*gorm.DB)
var initOnce = sync.Once{}

func Init(_ context.Context) error {
	var (
		err   error
		zetDB *gorm.DB
	)
	initOnce.Do(func() {
		zetDB, err = gormDB(conf.Conf().DBZet)
		if err != nil {
			return
		}
		dbMap["db_zet"] = zetDB
	})

	return err
}

func gormDB(c conf.Mysql) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(c.Dsn()), c.GormConfig())
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(c.MaxLifeTime))
	sqlDB.SetConnMaxIdleTime(time.Second * time.Duration(c.MaxIdleTime))
	return db, nil
}

func Zet() *gorm.DB {
	return dbMap["db_zet"]
}
